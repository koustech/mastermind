package server

import (
	"fmt"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"

	"github.com/aler9/gomavlib"
	"github.com/google/uuid"
	evbus "github.com/ispringtech/eventbus"
	pb "github.com/koustech/mastermind/gen/go/proto/mastermind/v1"

	u "github.com/koustech/mastermind/utils"
	"google.golang.org/grpc"
)

func Run(listenOn string, node *gomavlib.Node, sysId uint8, compId uint8) error {
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	gRPCServer := grpc.NewServer()
	serviceServer := NewMastermindServiceServer(node, sysId, compId)
	pb.RegisterMastermindServiceServer(gRPCServer, serviceServer)
	reflection.Register(gRPCServer)

	go GetTelem(serviceServer.stateBus, node)

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
	telemUpdateHandlers map[uuid.UUID]evbus.Subscription
	stateBus            evbus.Bus      // event notifier for state changes
	node                *gomavlib.Node // node connected to the plane
	sysId               uint8          // MAVLink system ID of the plane
	compId              uint8          // MAVLink component ID of the autopilot
}

func NewMastermindServiceServer(node *gomavlib.Node, sysId uint8, compId uint8) *mastermindServiceServer {
	// initializes state and allocates channels into empty channels slice
	stateBus := evbus.New()
	stateUpdateHandlers := make(map[uuid.UUID]evbus.Subscription) // handlers for every stateupdate func
	telemUpdateHandlers := make(map[uuid.UUID]evbus.Subscription) // handlers for every stateupdate func
	return &mastermindServiceServer{
		currentState:        pb.MissionState_MISSION_STATE_APPROACH,
		stateBus:            stateBus,
		stateUpdateHandlers: stateUpdateHandlers,
		telemUpdateHandlers: telemUpdateHandlers,
		node:                node,
		sysId:               sysId,
		compId:              compId,
	}
}
