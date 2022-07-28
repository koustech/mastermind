package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/aler9/gomavlib"
	"github.com/google/uuid"
	pb "github.com/koustech/mastermind/gen/go/proto/mastermind/v1"
	"github.com/koustech/mastermind/state"
	"github.com/koustech/mastermind/telemetry"
	u "github.com/koustech/mastermind/utils"
	"google.golang.org/grpc"
)

func main() {
	u.InitializeLoggers()
	defer u.SyncLogger()

	var grpcAddress string
	var mavlinkAddress string
	flag.StringVar(&grpcAddress, "a", "0.0.0.0:5678", "Mastermind's gRPC server address")
	flag.StringVar(&mavlinkAddress, "m", "127.0.0.1:14556", "The source MAVLink vehicle's ip address")
	flag.Parse()

	node := telemetry.ConnectToVehicle(mavlinkAddress)
	defer node.Close()

	if err := run(grpcAddress, node); err != nil {
		u.Logger.Fatal(err)
	}
}

func run(listenOn string, node *gomavlib.Node) error {
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}
	

	gRPCServer := grpc.NewServer()
	pb.RegisterMastermindServiceServer(gRPCServer, NewMastermindServiceServer())
	u.Logger.Infof("Listening on %v", listenOn)
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
	u.Logger.Info("new connection")

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
			u.Logger.Debugf("request received on session %v", sessionId)

			s.stateMu.Lock()
			oldState := s.currentState

			s.currentState, err = state.ResolveState(req.StateTransition, s.currentState)
			if err != nil {
				u.Logger.Warn(err)
			}

			u.Logger.Debug("calculated state")

			for service, ch := range s.sessions {
				u.Logger.Debugf("sending state response through %v channel...", service)
				ch <- &pb.UpdateStateResponse{OldState: oldState, StateTransition: req.StateTransition, CurrentState: s.currentState}
				u.Logger.Debugf("sent state response through %v channel", service)
			}
			s.stateMu.Unlock()
			u.Logger.Debugf("state mutex unlocked")
		}
	}

	cleanup := func() {
		u.Logger.Debugf("deleting channel for sessionId %v...", sessionId)
		close(s.sessions[sessionId])
		u.Logger.Debugf("deleted channel for sessionId %v", sessionId)

		u.Logger.Debugf("removing sessionId %v from sessions Map...", sessionId)
		delete(s.sessions, sessionId)
		u.Logger.Infof("removed sessionId %v from sessions Map", sessionId)
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
				u.Logger.Error("stream closed")
				return err
			}
		}
	}
}
