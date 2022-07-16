package main

import (
	"fmt"
	"github.com/koustech/mastermind/state"
	"io"
	"log"
	"net"
	"sync"

	pb "github.com/koustech/mastermind/gen/go/proto/mastermind/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterMastermindServiceServer(gRPCServer, NewMastermindServiceServer())
	log.Println("Listening on", listenOn)
	if err := gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

// mastermindServiceServer implements the MastermindService API.
type mastermindServiceServer struct {
	pb.UnimplementedMastermindServiceServer
	stateMu      sync.Mutex // protects currentState
	currentState pb.MissionState
	chansMu      sync.Mutex // protects chans
	chans        map[pb.SendingService]chan *pb.UpdateStateResponse
}

func NewMastermindServiceServer() *mastermindServiceServer {
	// initializes state and allocates channels into empty channels slice
	chans := make(map[pb.SendingService]chan *pb.UpdateStateResponse)
	return &mastermindServiceServer{
		currentState: pb.MissionState_MISSION_STATE_APPROACH,
		chans:        chans,
	}
}

// UpdateState updates the current state according to the state transition table
func (s *mastermindServiceServer) UpdateState(stream pb.MastermindService_UpdateStateServer) error {
	fmt.Println("new connection")
	channelChan := make(chan chan *pb.UpdateStateResponse) // holds the channel to be used throughout the stream. Only used once
	var sendingService pb.SendingService
	eof := make(chan bool)
	errChan := make(chan error)
	cleanup := func() {
		fmt.Println("deleting service key from channel")
		delete(s.chans, sendingService)
	}

	go func() {
		defer cleanup()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				eof <- true
				return
			}
			if err != nil {
				errChan <- err
				return
			}
			sendingService = req.SendingService
			fmt.Printf("request received from %v\n", req.SendingService)
			// create channel for sending service if doesn't exist
			_, exists := s.chans[req.SendingService]
			if !exists {
				//nameCh <- req.SendingService
				s.chans[req.SendingService] = make(chan *pb.UpdateStateResponse)

				fmt.Printf("created channel for %v\n", req.SendingService)

				channelChan <- s.chans[req.SendingService]
				fmt.Printf("transferred channel for %v\n", req.SendingService)
			} else {
				fmt.Println("channel already exists")
			}

			s.stateMu.Lock()
			oldState := s.currentState

			s.currentState, err = state.ResolveState(req.StateTransition, s.currentState)

			fmt.Println("calculated state")

			for service, ch := range s.chans {
				fmt.Printf("sending state response through %v channel...\n", service)
				ch <- &pb.UpdateStateResponse{OldState: oldState, StateTransition: req.StateTransition, CurrentState: s.currentState}
				fmt.Printf("sent state response through %v channel\n\n", service)
			}
			s.stateMu.Unlock()
		}
	}()

	fmt.Println("awaiting channel name...")
	var ch chan *pb.UpdateStateResponse
	ch = <-channelChan

	defer func() {
		fmt.Printf("closing %v, %v", channelChan, ch)
		cleanup()
		close(channelChan)
		close(ch)
	}()

	for {

		select {
		case <-eof:
			return nil

		case err := <-errChan:
			return err
		default:
		}

		if err := stream.Send(<-ch); err != nil {
			fmt.Println("stream closed")
			return err
		}
	}
}
