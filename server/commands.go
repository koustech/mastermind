package server

import (
	"context"
	"time"

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

// GotoWaypoint tells the plane to go to a specific point
func (s *mastermindServiceServer) GotoWaypoint(_ context.Context, request *pb.GotoWaypointRequest) (*pb.GotoWaypointResponse, error) {
	utils.Logger.Infof("GOTO lat:%v lon:%v", request.Lat, request.Lon)

	s.node.WriteMessageAll(&ardupilotmega.MessageMissionItem{
		TargetSystem:    s.sysId,
		TargetComponent: s.compId,
		Seq:             0,
		Frame:           ardupilotmega.MAV_FRAME_GLOBAL_RELATIVE_ALT,
		Command:         ardupilotmega.MAV_CMD_NAV_WAYPOINT,
		Current:         2,
		Autocontinue:    1,
		Param1:          0,
		Param2:          0,
		Param3:          0,
		Param4:          0,
		X:               request.Lat,
		Y:               request.Lon,
		Z:               request.Alt,
		MissionType: ardupilotmega.MAV_MISSION_TYPE_MISSION,
	})

	return &pb.GotoWaypointResponse{CommandSucceeded: true}, nil

}

const SET_ATTITUDE_TIMEOUT_SECONDS = 6

// SetAttitude tells the plane to obtain a specific attitude
func (s *mastermindServiceServer) SetAttitude(stream pb.MastermindService_SetAttitudeServer) error {
	// track time for case of timeout
	startTime := time.Now()
	utils.Logger.Info("SetAttitude() called")

	// set mode to FBWB ONE TIME ONLY

	//s.node.WriteMessageAll(&ardupilotmega.MessageCommandLong{
	//	TargetSystem:    s.sysId,
	//	TargetComponent: s.compId,
	//	Command:         ardupilotmega.MAV_CMD_DO_SET_MODE,
	//	Confirmation:    0,
	//	Param1:          float32(ardupilotmega.PLANE_MODE_FLY_BY_WIRE_A), // PARAM1 MUST BE MAIN MODE (FBWA)
	//	Param2:          float32(ardupilotmega.PLANE_MODE_FLY_BY_WIRE_B), // PARAM2 MUST BE SUBMODE (FBWB)
	//	Param3:          0,
	//	Param4:          0,
	//	Param5:          0,
	//	Param6:          0,
	//	Param7:          0,
	//})

	// TODO: implement
	// for every received message, send command and update target object (to be sent in detailed telem)

	for {
		currentTime := time.Now()
		if currentTime.Sub(startTime).Seconds() > 6 {
			break
		}

		attitudeRequest, err := stream.Recv()

		if err != nil {
			break
		}
		//s.node.WriteMessageAll(&ardupilotmega.MessageSetAttitudeTarget{
		//	TimeBootMs:      uint32(time.Now().UnixMilli()),
		//	TargetSystem:    s.sysId,
		//	TargetComponent: s.compId,
		//	TypeMask: ardupilotmega.ATTITUDE_TARGET_TYPEMASK_THROTTLE_IGNORE |
		//		ardupilotmega.ATTITUDE_TARGET_TYPEMASK_BODY_ROLL_RATE_IGNORE |
		//		ardupilotmega.ATTITUDE_TARGET_TYPEMASK_BODY_PITCH_RATE_IGNORE |
		//		ardupilotmega.ATTITUDE_TARGET_TYPEMASK_BODY_YAW_RATE_IGNORE,
		//	Q:             [4]float32{float32(quat.W), float32(quat.X), float32(quat.Y), float32(quat.Z)},
		//	BodyRollRate:  0,
		//	BodyPitchRate: 0,
		//	BodyYawRate:   0,
		//	Thrust:        0,
		//	ThrustBody:    [3]float32{},
		//})

		//utils.Logger.Info("roll = ch 1 = ", attitudeRequest.Roll)
		//utils.Logger.Info("pitch = ch 2 = ", attitudeRequest.Pitch)
		//s.node.WriteMessageAll(&ardupilotmega.MessageRcChannelsOverride{
		//	TargetSystem:    s.sysId,
		//	TargetComponent: s.compId,
		//	Chan1Raw:        1100, //roll
		//	Chan2Raw:        1500, //pitch
		//	Chan3Raw:        math.MaxUint16,
		//	Chan4Raw:        math.MaxUint16,
		//	Chan5Raw:        math.MaxUint16,
		//	Chan6Raw:        math.MaxUint16,
		//	Chan7Raw:        math.MaxUint16,
		//	Chan8Raw:        math.MaxUint16,
		//	Chan9Raw:        0,
		//	Chan10Raw:       0,
		//	Chan11Raw:       0,
		//	Chan12Raw:       0,
		//	Chan13Raw:       0,
		//	Chan14Raw:       0,
		//	Chan15Raw:       0,
		//	Chan16Raw:       0,
		//	Chan17Raw:       0,
		//	Chan18Raw:       0,
		//})

		//s.node.WriteMessageAll(&ardupilotmega.MessageCommandLong{
		//	TargetSystem:    s.sysId,
		//	TargetComponent: s.compId,
		//	Command:         ardupilotmega.MAV_CMD_DO_SET_SERVO,
		//	Confirmation:    0,
		//	Param1:          1,
		//	Param2:          attitudeRequest.Roll,
		//	Param3:          0,
		//	Param4:          0,
		//	Param5:          0,
		//	Param6:          0,
		//	Param7:          0,
		//})
		//s.node.WriteMessageAll(&ardupilotmega.MessageCommandLong{
		//	TargetSystem:    s.sysId,
		//	TargetComponent: s.compId,
		//	Command:         ardupilotmega.MAV_CMD_DO_SET_SERVO,
		//	Confirmation:    0,
		//	Param1:          2,
		//	Param2:          attitudeRequest.Pitch,
		//	Param3:          0,
		//	Param4:          0,
		//	Param5:          0,
		//	Param6:          0,
		//	Param7:          0,
		//})

		//s.node.WriteMessageAll(&ardupilotmega.MessageCommandLong{
		//	TargetSystem:    s.sysId,
		//	TargetComponent: s.compId,
		//	Command:         ardupilotmega.MAV_CMD_NAV_ATTITUDE_TIME,
		//	Confirmation:    0,
		//	Param1:          1,
		//	Param2:          attitudeRequest.Roll,
		//	Param3:          attitudeRequest.Pitch,
		//	Param4:          0,
		//	Param5:          3,
		//	Param6:          0,
		//	Param7:          0,
		//})

		s.target = attitudeRequest.Target
	}

	// close stream once EOF is received and set target to nil
	utils.Logger.Info("SetAttitude() finished")

	s.targetMu.Lock()
	s.target = nil
	s.targetMu.Unlock()

	return stream.SendAndClose(&pb.SetAttitudeResponse{})

}