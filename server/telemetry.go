package server

import (
	"time"

	"github.com/aler9/gomavlib"
	"github.com/aler9/gomavlib/pkg/dialects/ardupilotmega"
	"github.com/google/uuid"
	evbus "github.com/ispringtech/eventbus"
	pb "github.com/koustech/mastermind/gen/go/proto/mastermind/v1"
	u "github.com/koustech/mastermind/utils"
)

const (
	EventNewTelemetry = evbus.EventID("new_telemetry")
)

type NewTelemetryEvent struct {
	response *pb.GetDetailedTelemetryResponse
}

func (e *NewTelemetryEvent) EventID() evbus.EventID {
	return EventNewTelemetry
}

var Kamikazeflag = false
var kamikaze_start_time time.Time

const KamikazeSequence = 4
const WAIT_KAMIKAZE_TIME_S = 10

func GetTelem(s *mastermindServiceServer, node *gomavlib.Node) {
	vfrHud := ardupilotmega.MessageVfrHud{}
	attitude := ardupilotmega.MessageAttitude{}
	globalPositionInt := ardupilotmega.MessageGlobalPositionInt{}
	navControllerOutput := ardupilotmega.MessageNavControllerOutput{}
	batteryStatus := ardupilotmega.MessageBatteryStatus{}
	wind := ardupilotmega.MessageWind{}

	isAutonomous := false // plane is autonomous if not in FBWA, MANUAL, or AUTOTUNE modes

	var timeBootMs uint64

	var updateFlag bool
	//kamikaze_wait := time.Now()

	// print every message we receive
	for evt := range node.Events() {

		updateFlag = true
		if frm, ok := evt.(*gomavlib.EventFrame); ok {
			if frm.SystemID() != 1 || frm.ComponentID() != 1 {
				continue
			}
			switch msg := frm.Message().(type) {
			case *ardupilotmega.MessageAttitude:
				attitude = *msg
			case *ardupilotmega.MessageVfrHud:
				vfrHud = *msg
			case *ardupilotmega.MessageGlobalPositionInt:
				globalPositionInt = *msg
			case *ardupilotmega.MessageNavControllerOutput:
				navControllerOutput = *msg
			case *ardupilotmega.MessageBatteryStatus:
				batteryStatus = *msg
				updateFlag = false // Don't want to increase message sending frequency just for battery
			case *ardupilotmega.MessageSystemTime:
				timeBootMs = msg.TimeUnixUsec / 1000
			case *ardupilotmega.MessageMissionAck:
				u.Logger.Debugf("received mission ack: %+v", msg)
				updateFlag = false
			case *ardupilotmega.MessageWind:
				wind = *msg
				updateFlag = false
			case *ardupilotmega.MessageHeartbeat:
				switch ardupilotmega.PLANE_MODE(msg.CustomMode) {
				case ardupilotmega.PLANE_MODE_MANUAL,
					ardupilotmega.PLANE_MODE_FLY_BY_WIRE_A,
					ardupilotmega.PLANE_MODE_FLY_BY_WIRE_B,
					ardupilotmega.PLANE_MODE_AUTOTUNE:

					isAutonomous = false
				default:
					isAutonomous = true
				}

			case *ardupilotmega.MessageMissionItemReached:
				u.Logger.Debugf("received mission item reached: %+v", msg)
				recieved := msg.Seq
				if recieved == KamikazeSequence {
					if s.currentState == pb.MissionState_MISSION_STATE_KAMIKAZE {
						kamikaze_start_time = time.Now()
						u.Logger.Warnf("Kamikaze waypoint reached UCHAK DALIOGHR!!!")
						u.Logger.Warnf("kamikaze started at: %v", kamikaze_start_time)
						SetModeKamikaze(frm.SystemID(), frm.ComponentID(), node)
					}
				}

			// case *ardupilotmega.MessageMissionCurrent:

			// 	recieved := msg.Seq
			// 	if recieved == KamikazeSequence {
			// 		elapsed_kamikaze_time := time.Now()
			// 		if elapsed_kamikaze_time.Sub(kamikaze_wait) > time.Duration(WAIT_KAMIKAZE_TIME_S)*time.Second {
			// 			Kamikazeflag = false
			// 		}
			// 		if s.currentState == pb.MissionState_MISSION_STATE_KAMIKAZE && !Kamikazeflag {
			// 			u.Logger.Warnf("Kamikaze waypoint reached UCHAK DALIOGHR!!!")
			// 			kamikaze_wait = time.Now()
			// 			Kamikazeflag = true
			// 			SetModeKamikaze(frm.SystemID(), frm.ComponentID(), node)
			// 		}
			// 	}
			default:
				updateFlag = false
			}

			if updateFlag {
				s.stateBus.Publish(&NewTelemetryEvent{
					response: &pb.GetDetailedTelemetryResponse{
						TimeBootMs:    timeBootMs,
						Lat:           globalPositionInt.Lat,
						Lon:           globalPositionInt.Lon,
						RelativeAlt:   globalPositionInt.RelativeAlt,
						Roll:          attitude.Roll,
						Pitch:         attitude.Pitch,
						Yaw:           attitude.Yaw,
						Airspeed:      vfrHud.Airspeed,
						Groundspeed:   vfrHud.Groundspeed,
						WpDist:        uint32(navControllerOutput.WpDist),
						Battery:       int32(batteryStatus.BatteryRemaining),
						Autonomous:    isAutonomous,
						Target:        nil,
						WindSpeed:     wind.Speed,
						WindDirection: wind.Direction,
					},
				})
			}
		}
	}
}

// GetTelemetry sends a stream of telemetry to all clients
func (s *mastermindServiceServer) GetTelemetry(_ *pb.GetTelemetryRequest, stream pb.MastermindService_GetTelemetryServer) error {
	// generate new sessionId and add channel
	sessionId := uuid.New()
	u.Logger.Info("new connection for session: ", sessionId)

	errChan := make(chan error)

	s.telemUpdateHandlers[sessionId] = s.stateBus.Subscribe(EventNewTelemetry, func(e evbus.Event) {
		se := e.(*NewTelemetryEvent)

		// change detailed telemetry to normal telem response
		response := &pb.GetTelemetryResponse{
			TimeBootMs:    se.response.TimeBootMs,
			Lat:           se.response.Lat,
			Lon:           se.response.Lon,
			RelativeAlt:   se.response.RelativeAlt,
			Roll:          se.response.Roll,
			Pitch:         se.response.Pitch,
			Yaw:           se.response.Yaw,
			Airspeed:      se.response.Airspeed,
			Groundspeed:   se.response.Groundspeed,
			WpDist:        se.response.WpDist,
			WindSpeed:     se.response.WindSpeed,
			WindDirection: se.response.WindDirection,
		}
		if err := stream.Send(response); err != nil {
			s.stateBus.Unsubscribe(s.telemUpdateHandlers[sessionId])
			delete(s.telemUpdateHandlers, sessionId)
			u.Logger.Infof("unsubscribed telem session: %v from event bus", sessionId)

			errChan <- err
		}
	})

	<-errChan
	return nil
}

// GetDetailedTelemetry sends a stream of detailed telemetry to all clients
func (s *mastermindServiceServer) GetDetailedTelemetry(_ *pb.GetDetailedTelemetryRequest, stream pb.MastermindService_GetDetailedTelemetryServer) error {
	// generate new sessionId and add channel
	sessionId := uuid.New()
	u.Logger.Info("new connection for session: ", sessionId)

	errChan := make(chan error)

	s.telemUpdateHandlers[sessionId] = s.stateBus.Subscribe(EventNewTelemetry, func(e evbus.Event) {
		se := e.(*NewTelemetryEvent)
		response := se.response
		s.targetMu.Lock()
		se.response.Target = s.target
		s.targetMu.Unlock()

		if err := stream.Send(response); err != nil {
			s.stateBus.Unsubscribe(s.telemUpdateHandlers[sessionId])
			delete(s.telemUpdateHandlers, sessionId)
			u.Logger.Infof("unsubscribed telem session: %v from event bus", sessionId)

			errChan <- err
		}
	})

	<-errChan
	return nil
}
