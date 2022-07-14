package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	pb "github.com/koustech/mastermind/gen/proto/go/mastermind/v1"
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
	mu            sync.Mutex // protects current_state
	current_state pb.MissionState
}

// UpdateState updates the current state according to the state transition table
func (s *mastermindServiceServer) UpdateState(stream pb.MastermindService_UpdateStateServer) error {
	s.current_state = pb.MissionState_MISSION_STATE_APPROACH
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		s.mu.Lock()

		switch s.current_state {
		case pb.MissionState_MISSION_STATE_APPROACH:
			stream.Send(&pb.UpdateStateResponse{CurrentState: pb.MissionState_MISSION_STATE_APPROACH})
		case pb.MissionState_MISSION_STATE_FOLLOWING:
			stream.Send(&pb.UpdateStateResponse{CurrentState: pb.MissionState_MISSION_STATE_FOLLOWING})
		case pb.MissionState_MISSION_STATE_KAMIKAZE:
			stream.Send(&pb.UpdateStateResponse{CurrentState: pb.MissionState_MISSION_STATE_KAMIKAZE})
		default:
			stream.Send(&pb.UpdateStateResponse{CurrentState: pb.MissionState_MISSION_STATE_APPROACH})
		}
		s.mu.Unlock()
	}
}
