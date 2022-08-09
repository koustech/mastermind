package server

import (
	"context"
	"github.com/aler9/gomavlib/pkg/dialects/ardupilotmega"
	pb "github.com/koustech/mastermind/gen/go/proto/mastermind/v1"
	"github.com/koustech/mastermind/utils"
)

// SetPIDLevel sets the PID of the plane according to predetermined levels in the config
func (s *mastermindServiceServer) SetPIDLevel(_ context.Context, request *pb.SetPIDLevelRequest) (*pb.SetPIDLevelResponse, error) {
	// FIXME: add actual command invocation
	utils.Logger.Infof("PID Level %v request received", request.Level)
	success := true

	switch request.Level {
	case pb.PIDLevel_PID_LEVEL_1:
		utils.Logger.Info("PID set to ", request.Level)
	case pb.PIDLevel_PID_LEVEL_2:
		utils.Logger.Info("PID set to ", request.Level)
	case pb.PIDLevel_PID_LEVEL_3:
		utils.Logger.Info("PID set to ", request.Level)
	case pb.PIDLevel_PID_LEVEL_4:
		utils.Logger.Info("PID set to ", request.Level)
	case pb.PIDLevel_PID_LEVEL_5:
		utils.Logger.Info("PID set to ", request.Level)
	case pb.PIDLevel_PID_LEVEL_6:
		utils.Logger.Info("PID set to ", request.Level)
	case pb.PIDLevel_PID_LEVEL_7:
		utils.Logger.Info("PID set to ", request.Level)
	case pb.PIDLevel_PID_LEVEL_8:
		utils.Logger.Info("PID set to ", request.Level)
	default:
		utils.Logger.Errorf("level %v not supported", request.Level)
		success = false
	}

	return &pb.SetPIDLevelResponse{CommandSucceeded: success}, nil
}

// SetSpeed sets the airspeed of the plane
func (s *mastermindServiceServer) SetSpeed(_ context.Context, request *pb.SetSpeedRequest) (*pb.SetSpeedResponse, error) {
	utils.Logger.Infof("Setting speed to %v m/s", request.NewSpeed)

	s.node.WriteMessageAll(&ardupilotmega.MessageCommandLong{
		TargetSystem:    s.sysId,
		TargetComponent: s.compId,
		Command:         ardupilotmega.MAV_CMD_DO_CHANGE_SPEED,
		Confirmation:    0,                // TODO: should this be 0 or 1?
		Param1:          0,                // Speed Type --> airspeed=0, groundspeed=1, climbspeed=2
		Param2:          request.NewSpeed, // Speed
		Param3:          0,                // Throttle
		Param4:          0,
		Param5:          0,
		Param6:          0,
		Param7:          0,
	})
	utils.Logger.Infof("Set speed to %v m/s", request.NewSpeed)

	return &pb.SetSpeedResponse{CommandSucceeded: true}, nil
}
