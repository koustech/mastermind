package server

import (
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
	response *pb.GetTelemetryResponse
}

func (e *NewTelemetryEvent) EventID() evbus.EventID {
	return EventNewTelemetry
}

func GetTelem(bus evbus.Bus, node *gomavlib.Node) {
	vfrHud := ardupilotmega.MessageVfrHud{}
	attitude := ardupilotmega.MessageAttitude{}
	globalPositionInt := ardupilotmega.MessageGlobalPositionInt{}
	navControllerOutput := ardupilotmega.MessageNavControllerOutput{}
	var timeBootMs uint32

	var updateFlag bool
	// print every message we receive
	for evt := range node.Events() {
		updateFlag = true

		if frm, ok := evt.(*gomavlib.EventFrame); ok {
			switch msg := frm.Message().(type) {
			case *ardupilotmega.MessageAttitude:
				attitude = *msg
				timeBootMs = msg.TimeBootMs
			case *ardupilotmega.MessageVfrHud:
				vfrHud = *msg
			case *ardupilotmega.MessageGlobalPositionInt:
				globalPositionInt = *msg
				timeBootMs = msg.TimeBootMs
			case *ardupilotmega.MessageNavControllerOutput:
				navControllerOutput = *msg
				u.Logger.Debug(navControllerOutput)
			default:
				updateFlag = false
			}

			if updateFlag {
				bus.Publish(&NewTelemetryEvent{
					response: &pb.GetTelemetryResponse{
						TimeBootMs:  timeBootMs,
						Lat:         globalPositionInt.Lat,
						Lon:         globalPositionInt.Lon,
						RelativeAlt: globalPositionInt.RelativeAlt,
						Roll:        attitude.Roll,
						Pitch:       attitude.Pitch,
						Yaw:         attitude.Yaw,
						Airspeed:    vfrHud.Airspeed,
						Groundspeed: vfrHud.Groundspeed,
						WpDist:      uint32(navControllerOutput.WpDist),
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
		response := se.response
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
