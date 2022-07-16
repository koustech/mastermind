package state

import (
	"errors"

	. "github.com/koustech/mastermind/gen/go/proto/mastermind/v1"
)

// ErrInvalidStateTransition is the error returned by ResolveState when an invalid StateTransition happens
var ErrInvalidStateTransition = errors.New("current state has no defined behavior for requested state transition")

// ResolveState resolves the MissionState according to the current MissionState and StateTransition given as input
func ResolveState(state_transition StateTransition, current_state MissionState) (MissionState, error) {
	switch current_state {
	case MissionState_MISSION_STATE_APPROACH:
		switch state_transition {
		case StateTransition_STATE_TRANSITION_TARGET_APPROACHED:
			return MissionState_MISSION_STATE_FOLLOWING, nil
		case StateTransition_STATE_TRANSITION_MODE_KAMIKAZE_SELECTED:
			return MissionState_MISSION_STATE_KAMIKAZE, nil
		default:
			return MissionState_MISSION_STATE_APPROACH, ErrInvalidStateTransition
		}
	case MissionState_MISSION_STATE_FOLLOWING:
		switch state_transition {
		case StateTransition_STATE_TRANSITION_LOCK_FAILED:
			return MissionState_MISSION_STATE_APPROACH, nil
		case StateTransition_STATE_TRANSITION_LOCK_SUCCESS:
			return MissionState_MISSION_STATE_APPROACH, nil
		case StateTransition_STATE_TRANSITION_MODE_KAMIKAZE_SELECTED:
			return MissionState_MISSION_STATE_KAMIKAZE, nil
		default:
			return MissionState_MISSION_STATE_FOLLOWING, ErrInvalidStateTransition
		}
	case MissionState_MISSION_STATE_KAMIKAZE:
		switch state_transition {
		case StateTransition_STATE_TRANSITION_QR_FAILED:
			return MissionState_MISSION_STATE_KAMIKAZE, nil
		case StateTransition_STATE_TRANSITION_QR_SUCCESS:
			return MissionState_MISSION_STATE_APPROACH, nil
		case StateTransition_STATE_TRANSITION_MODE_APPROACH_SELECTED:
			return MissionState_MISSION_STATE_APPROACH, nil
		default:
			return MissionState_MISSION_STATE_KAMIKAZE, ErrInvalidStateTransition
		}
	default:
		return MissionState_MISSION_STATE_APPROACH, ErrInvalidStateTransition
	}
}
