package server

import (
	"context"
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
