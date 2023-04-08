package state

import (
	"testing"

	. "github.com/koustech/mastermind/gen/go/proto/mastermind/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestResolveStateValidTransitions checks that transitions are correctly resolved for all defined state transitions
func TestResolveStateValidTransitions(t *testing.T) {
	// Setup
	testCases := []struct {
		desc            string
		stateTransition StateTransition
		inputState      MissionState
		resultState     MissionState
		err             error
	}{
		{
			desc:            "transition from approach state to detecting state with target_approached",
			stateTransition: StateTransition_STATE_TRANSITION_TARGET_APPROACHED,
			inputState:      MissionState_MISSION_STATE_APPROACH,
			resultState:     MissionState_MISSION_STATE_DETECTING,
			err:             nil,
		},
		{
			desc:            "transition from approach state to detecting_gps state with target_approached_gps",
			stateTransition: StateTransition_STATE_TRANSITION_TARGET_APPROACHED_GPS,
			inputState:      MissionState_MISSION_STATE_APPROACH,
			resultState:     MissionState_MISSION_STATE_DETECTING_GPS,
			err:             nil,
		},
		{
			desc:            "transition from approach state to kamikaze state with mode kamikaze selected",
			stateTransition: StateTransition_STATE_TRANSITION_MODE_KAMIKAZE_SELECTED,
			inputState:      MissionState_MISSION_STATE_APPROACH,
			resultState:     MissionState_MISSION_STATE_KAMIKAZE,
			err:             nil,
		},
		{
			desc:            "transition from detecting_gps state to lock_gps state with target_detected",
			stateTransition: StateTransition_STATE_TRANSITION_TARGET_DETECTED,
			inputState:      MissionState_MISSION_STATE_DETECTING_GPS,
			resultState:     MissionState_MISSION_STATE_LOCK_GPS,
			err:             nil,
		},
		{
			desc:            "transition from detecting state to lock state with target_detected",
			stateTransition: StateTransition_STATE_TRANSITION_TARGET_DETECTED,
			inputState:      MissionState_MISSION_STATE_DETECTING,
			resultState:     MissionState_MISSION_STATE_LOCK,
			err:             nil,
		},
		{
			desc:            "transition from lock state to approach state with lock_failed",
			stateTransition: StateTransition_STATE_TRANSITION_LOCK_FAILED,
			inputState:      MissionState_MISSION_STATE_LOCK,
			resultState:     MissionState_MISSION_STATE_APPROACH,
			err:             nil,
		},
		{
			desc:            "transition from lock_gps state to approach state with lock_failed",
			stateTransition: StateTransition_STATE_TRANSITION_LOCK_FAILED,
			inputState:      MissionState_MISSION_STATE_LOCK_GPS,
			resultState:     MissionState_MISSION_STATE_APPROACH,
			err:             nil,
		},
		{
			desc:            "transition from lock state to approach state with lock_success",
			stateTransition: StateTransition_STATE_TRANSITION_LOCK_SUCCESS,
			inputState:      MissionState_MISSION_STATE_LOCK,
			resultState:     MissionState_MISSION_STATE_APPROACH,
			err:             nil,
		},
		{
			desc:            "transition from lock_gps state to approach state with lock_success",
			stateTransition: StateTransition_STATE_TRANSITION_LOCK_SUCCESS,
			inputState:      MissionState_MISSION_STATE_LOCK_GPS,
			resultState:     MissionState_MISSION_STATE_APPROACH,
			err:             nil,
		},
		{
			desc:            "transition from lock state to kamikaze state with mode kamikaze selected",
			stateTransition: StateTransition_STATE_TRANSITION_MODE_KAMIKAZE_SELECTED,
			inputState:      MissionState_MISSION_STATE_LOCK,
			resultState:     MissionState_MISSION_STATE_KAMIKAZE,
			err:             nil,
		},
		{
			desc:            "transition from lock_gps state to kamikaze state with mode kamikaze selected",
			stateTransition: StateTransition_STATE_TRANSITION_MODE_KAMIKAZE_SELECTED,
			inputState:      MissionState_MISSION_STATE_LOCK_GPS,
			resultState:     MissionState_MISSION_STATE_KAMIKAZE,
			err:             nil,
		},
		{
			desc:            "transition from kamikaze state to kamikaze state with QR failed",
			stateTransition: StateTransition_STATE_TRANSITION_QR_FAILED,
			inputState:      MissionState_MISSION_STATE_KAMIKAZE,
			resultState:     MissionState_MISSION_STATE_KAMIKAZE,
			err:             nil,
		},
		{
			desc:            "transition from kamikaze state to approach state with QR success",
			stateTransition: StateTransition_STATE_TRANSITION_QR_SUCCESS,
			inputState:      MissionState_MISSION_STATE_KAMIKAZE,
			resultState:     MissionState_MISSION_STATE_APPROACH,
			err:             nil,
		},
		{
			desc:            "transition from kamikaze state to approach state with mode approach selected",
			stateTransition: StateTransition_STATE_TRANSITION_MODE_APPROACH_SELECTED,
			inputState:      MissionState_MISSION_STATE_KAMIKAZE,
			resultState:     MissionState_MISSION_STATE_APPROACH,
			err:             nil,
		},
	}

	for _, scenario := range testCases {
		t.Run(scenario.desc, func(t *testing.T) {
			result, err := ResolveState(scenario.stateTransition, scenario.inputState)
			require.NoError(t, err)
			assert.Equal(t, scenario.resultState, result)
		})
	}
}
