package main

import (
	"fmt"
	"github.com/google/uuid"
	pb "github.com/koustech/mastermind/gen/go/proto/mastermind/v1"
	"github.com/koustech/mastermind/state"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"sync"
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
	chansMu      sync.Mutex                                 // protects sessions
	sessions     map[uuid.UUID]chan *pb.UpdateStateResponse // all active sessions
}

func NewMastermindServiceServer() *mastermindServiceServer {
	// initializes state and allocates channels into empty channels slice
	chans := make(map[uuid.UUID]chan *pb.UpdateStateResponse)
	return &mastermindServiceServer{
		currentState: pb.MissionState_MISSION_STATE_APPROACH,
		sessions:     chans,
	}
}

// UpdateState updates the current state according to the state transition table
func (s *mastermindServiceServer) UpdateState(stream pb.MastermindService_UpdateStateServer) error {

	fmt.Println("new connection")

	// generate new sessionId and add channel
	sessionId := uuid.New()
	s.sessions[sessionId] = make(chan *pb.UpdateStateResponse)

	eof := make(chan bool)
	errChan := make(chan error)
	receiverThread := func() {
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
			fmt.Printf("request received on session %v\n", sessionId)

			s.stateMu.Lock()
			oldState := s.currentState

			s.currentState, err = state.ResolveState(req.StateTransition, s.currentState)

			fmt.Println("calculated state")

			for service, ch := range s.sessions {
				fmt.Printf("sending state response through %v channel...\n", service)
				ch <- &pb.UpdateStateResponse{OldState: oldState, StateTransition: req.StateTransition, CurrentState: s.currentState}
				fmt.Printf("sent state response through %v channel\n\n", service)
			}
			s.stateMu.Unlock()
		}
	}

	cleanup := func() {
		fmt.Println("deleting channel...")
		close(s.sessions[sessionId])
		fmt.Println("removing sessionId from sessions list...")

		delete(s.sessions, sessionId)

	}

	go receiverThread()

	defer cleanup()

	for {
		select {
		case <-eof:
			return nil
		case err := <-errChan:
			return err
		case response := <-s.sessions[sessionId]:
			if err := stream.Send(response); err != nil {
				fmt.Println("stream closed")
				return err
			}
		}
	}
}
