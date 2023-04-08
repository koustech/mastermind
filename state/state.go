package state

import (
	"errors"

	. "github.com/koustech/mastermind/gen/go/proto/mastermind/v1"
)

// ErrInvalidStateTransition is the error returned by ResolveState when an invalid StateTransition happens
var ErrInvalidStateTransition = errors.New("current state has no defined behavior for requested state transition")

// ResolveState resolves the MissionState according to the current MissionState and StateTransition given as input
func ResolveState(stateTransition StateTransition, currentState MissionState) (MissionState, error) {

	// mode overrides
	switch stateTransition {
	case StateTransition_STATE_TRANSITION_MODE_KAMIKAZE_SELECTED:
		return MissionState_MISSION_STATE_KAMIKAZE, nil
	case StateTransition_STATE_TRANSITION_MODE_APPROACH_SELECTED:
		return MissionState_MISSION_STATE_APPROACH, nil
	case StateTransition_STATE_TRANSITION_MODE_NEUTRAL_SELECTED:
		return MissionState_MISSION_STATE_NEUTRAL, nil
	}

	switch currentState {
	case MissionState_MISSION_STATE_APPROACH:
		switch stateTransition {
		case StateTransition_STATE_TRANSITION_TARGET_APPROACHED:
			return MissionState_MISSION_STATE_DETECTING, nil
		case StateTransition_STATE_TRANSITION_TARGET_APPROACHED_GPS:
			return MissionState_MISSION_STATE_DETECTING_GPS, nil
		case StateTransition_STATE_TRANSITION_MODE_KAMIKAZE_SELECTED:
			return MissionState_MISSION_STATE_KAMIKAZE, nil
		default:
			return MissionState_MISSION_STATE_APPROACH, ErrInvalidStateTransition
		}
	case MissionState_MISSION_STATE_DETECTING:
		switch stateTransition {
		case StateTransition_STATE_TRANSITION_TARGET_DETECTED:
			return MissionState_MISSION_STATE_LOCK, nil
		default:
			return MissionState_MISSION_STATE_DETECTING, ErrInvalidStateTransition
		}
	case MissionState_MISSION_STATE_DETECTING_GPS:
		switch stateTransition {
		case StateTransition_STATE_TRANSITION_TARGET_DETECTED:
			return MissionState_MISSION_STATE_LOCK_GPS, nil
		default:
			return MissionState_MISSION_STATE_DETECTING_GPS, ErrInvalidStateTransition
		}
	case MissionState_MISSION_STATE_LOCK:
		switch stateTransition {
		case StateTransition_STATE_TRANSITION_LOCK_FAILED:
			return MissionState_MISSION_STATE_APPROACH, nil
		case StateTransition_STATE_TRANSITION_LOCK_SUCCESS:
			return MissionState_MISSION_STATE_APPROACH, nil
		case StateTransition_STATE_TRANSITION_MODE_KAMIKAZE_SELECTED:
			return MissionState_MISSION_STATE_KAMIKAZE, nil
		default:
			return MissionState_MISSION_STATE_LOCK, ErrInvalidStateTransition
		}
	case MissionState_MISSION_STATE_LOCK_GPS:
		switch stateTransition {
		case StateTransition_STATE_TRANSITION_LOCK_FAILED:
			return MissionState_MISSION_STATE_APPROACH, nil
		case StateTransition_STATE_TRANSITION_LOCK_SUCCESS:
			return MissionState_MISSION_STATE_APPROACH, nil
		case StateTransition_STATE_TRANSITION_MODE_KAMIKAZE_SELECTED:
			return MissionState_MISSION_STATE_KAMIKAZE, nil
		default:
			return MissionState_MISSION_STATE_LOCK_GPS, ErrInvalidStateTransition
		}
	case MissionState_MISSION_STATE_KAMIKAZE:
		switch stateTransition {
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
