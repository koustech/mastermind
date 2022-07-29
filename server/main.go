package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/aler9/gomavlib"
	"github.com/google/uuid"
	evbus "github.com/ispringtech/eventbus"
	pb "github.com/koustech/mastermind/gen/go/proto/mastermind/v1"
	"github.com/koustech/mastermind/state"

	// "github.com/koustech/mastermind/telemetry"
	u "github.com/koustech/mastermind/utils"
	"google.golang.org/grpc"
)

const (
	eventNewState = evbus.EventID("new_state")
)

type newStateEvent struct {
	response *pb.UpdateStateResponse
}

func (e *newStateEvent) EventID() evbus.EventID {
	return eventNewState
}

func main() {
	u.InitializeLoggers()
	defer u.SyncLogger()

	var grpcAddress string
	var mavlinkAddress string
	flag.StringVar(&grpcAddress, "a", "0.0.0.0:5678", "Mastermind's gRPC server address")
	flag.StringVar(&mavlinkAddress, "m", "127.0.0.1:14556", "The source MAVLink vehicle's ip address")
	flag.Parse()

	// node := telemetry.ConnectToVehicle(mavlinkAddress)
	node := &gomavlib.Node{}
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

	stateUpdateHandlers map[uuid.UUID]evbus.Subscription // stateUpdateFuncs for all active sessison
	stateBus            evbus.Bus                        // event notifier for state changes
}

func NewMastermindServiceServer() *mastermindServiceServer {
	// initializes state and allocates channels into empty channels slice
	stateBus := evbus.New()
	stateUpdateHandlers := make(map[uuid.UUID]evbus.Subscription) // handlers for every stateupdate func
	return &mastermindServiceServer{
		currentState:        pb.MissionState_MISSION_STATE_APPROACH,
		stateBus:            stateBus,
		stateUpdateHandlers: stateUpdateHandlers,
	}
}

// UpdateState updates the current state according to the state transition table
func (s *mastermindServiceServer) UpdateState(stream pb.MastermindService_UpdateStateServer) error {
	// generate new sessionId and add channel
	sessionId := uuid.New()
	u.Logger.Info("new connection for session: ", sessionId)

	errChan := make(chan error)

	cleanup := func() {
		s.stateBus.Unsubscribe(s.stateUpdateHandlers[sessionId])
		delete(s.stateUpdateHandlers, sessionId)
		u.Logger.Infof("unsubscribed session: %v from event bus", sessionId)
	}

	defer cleanup()
	s.stateUpdateHandlers[sessionId] = s.stateBus.Subscribe(eventNewState, func(e evbus.Event) {
		se := e.(*newStateEvent)
		response := se.response
		if err := stream.Send(response); err != nil {
			u.Logger.Error("stream closed")
			errChan <- err
		}
		u.Logger.Debugf("sent state response through %v channel", sessionId)
	})

	for {
		select {
		case err := <-errChan:
			u.Logger.Errorf("error received: %v", err)
			return err
		default:
		}

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		u.Logger.Debugf("request received on session %v", sessionId)

		s.stateMu.Lock()
		u.Logger.Debugf("state mutex locked")

		oldState := s.currentState

		s.currentState, err = state.ResolveState(req.StateTransition, s.currentState)
		if err != nil {
			u.Logger.Warn(err)
		}
		u.Logger.Infof("%v -> %v", oldState, s.currentState)

		s.stateBus.Publish(&newStateEvent{response: &pb.UpdateStateResponse{OldState: oldState, StateTransition: req.StateTransition, CurrentState: s.currentState}})

		s.stateMu.Unlock()
		u.Logger.Debugf("state mutex unlocked")
	}
}
