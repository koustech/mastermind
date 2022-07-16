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
	mastermindServer := NewMastermindServiceServer()
	pb.RegisterMastermindServiceServer(gRPCServer, mastermindServer)

	go mastermindServer.SendResponses()
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
	streams      map[uuid.UUID]pb.MastermindService_UpdateStateServer
	allSent      chan struct{}
}

func NewMastermindServiceServer() *mastermindServiceServer {
	// initializes state and allocates channels into empty channels slice
	chans := make(map[uuid.UUID]chan *pb.UpdateStateResponse)
	streams := make(map[uuid.UUID]pb.MastermindService_UpdateStateServer)
	allSent := make(chan struct{}) // sends when all clients are notified
	return &mastermindServiceServer{
		currentState: pb.MissionState_MISSION_STATE_APPROACH,
		sessions:     chans,
		streams:      streams,
		allSent:      allSent,
	}
}

//SendResponses checks for new data and sends responses to streams accordingly
func (s *mastermindServiceServer) SendResponses() {
	for {
		wg := new(sync.WaitGroup)
		if len(s.streams) == 0 {
			continue
		}
		wg.Add(len(s.streams))
		for sessionId, stateServer := range s.streams {
			go func(stateServer pb.MastermindService_UpdateStateServer, sessionId uuid.UUID) {
				err := stateServer.Send(<-s.sessions[sessionId])
				if err != nil {
					fmt.Println("stream closed")
				}
				wg.Done()
			}(stateServer, sessionId)
		}
		wg.Wait()
		s.allSent <- struct{}{}
	}
}

// UpdateState updates the current state according to the state transition table
func (s *mastermindServiceServer) UpdateState(stream pb.MastermindService_UpdateStateServer) error {

	fmt.Println("new connection")

	// generate new sessionId and add channel
	sessionId := uuid.New()
	s.sessions[sessionId] = make(chan *pb.UpdateStateResponse)

	// add stream to streams map
	s.streams[sessionId] = stream

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
			fmt.Println("mutex locked")
			oldState := s.currentState

			s.currentState, err = state.ResolveState(req.StateTransition, s.currentState)

			fmt.Println("calculated state")

			for service, ch := range s.sessions {
				// send to all channels concurrently
				go func(service uuid.UUID, ch chan *pb.UpdateStateResponse) {
					fmt.Printf("sending state response through %v channel...\n", service)
					ch <- &pb.UpdateStateResponse{OldState: oldState, StateTransition: req.StateTransition, CurrentState: s.currentState}
					fmt.Printf("sent state response through %v channel\n", service)
				}(service, ch)
			}
		}
	}

	cleanup := func() {
		s.stateMu.Lock()
		fmt.Println("deleting channel...")
		close(s.sessions[sessionId])
		fmt.Println("removing sessionId from sessions list...")

		delete(s.sessions, sessionId)

		fmt.Println("deleting stream")

		delete(s.streams, sessionId)
		// FIXME: UNLOCK OF UNLOCKED MUTEX HERE
		s.stateMu.Unlock()
	}
	defer cleanup()

	go receiverThread()

	for {
		select {
		case <-eof:
			return nil
		case err := <-errChan:
			return err
		case <-s.allSent:
			// unlock mutex only after update is sent to all clients
			s.stateMu.Unlock()
			fmt.Println("mutex unlocked")
		}
	}
}
