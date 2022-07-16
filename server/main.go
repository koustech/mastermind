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

	server := grpc.NewServer()
	pb.RegisterMastermindServiceServer(server, &mastermindServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

// mastermindServiceServer implements the MastermindService API.
type mastermindServiceServer struct {
	pb.UnimplementedMastermindServiceServer
	mu           sync.Mutex // protects currentState
	currentState pb.MissionState
}

// UpdateState updates the current state according to the state transition table
func (s *mastermindServiceServer) UpdateState(stream pb.MastermindService_UpdateStateServer) error {
	s.currentState = pb.MissionState_MISSION_STATE_APPROACH
	for {
		// durum guncellemesini al
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		s.mu.Lock()
		oldState := s.currentState
		// yeni durumu hesapla ve yeni durumu guncelle
		s.currentState, err = state.ResolveState(req.StateTransition, s.currentState)
		s.mu.Unlock()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		// yeni durumu istemciye gonder
		if err := stream.Send(&pb.UpdateStateResponse{OldState: oldState, StateTransition: req.StateTransition, CurrentState: s.currentState}); err != nil {
			return err
		}
	}
}
