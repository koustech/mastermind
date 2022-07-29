package server

import (
	"io"

	"github.com/google/uuid"
	evbus "github.com/ispringtech/eventbus"
	pb "github.com/koustech/mastermind/gen/go/proto/mastermind/v1"
	"github.com/koustech/mastermind/state"

	u "github.com/koustech/mastermind/utils"
)

const (
	EventNewState = evbus.EventID("new_state")
)

type NewStateEvent struct {
	response *pb.UpdateStateResponse
}

func (e *NewStateEvent) EventID() evbus.EventID {
	return EventNewState
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
	s.stateUpdateHandlers[sessionId] = s.stateBus.Subscribe(EventNewState, func(e evbus.Event) {
		se := e.(*NewStateEvent)
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

		s.stateBus.Publish(&NewStateEvent{response: &pb.UpdateStateResponse{OldState: oldState, StateTransition: req.StateTransition, CurrentState: s.currentState}})

		s.stateMu.Unlock()
		u.Logger.Debugf("state mutex unlocked")
	}
}
