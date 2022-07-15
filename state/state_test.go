package state

import (
	"testing"

	. "github.com/koustech/mastermind/gen/proto/go/mastermind/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestResolveStateErrors checks that input states with undefined transitions
// return an error and the same state
func TestResolveStateErrors(t *testing.T) {
	// Setup
	testCases := []struct {
		desc            string
		stateTransition StateTransition
		inputState      MissionState
		resultState     MissionState
		err             error
	}{
		{
			desc:            "transition from approach state with approach_selected = ERR",
			stateTransition: StateTransition_STATE_TRANSITION_MODE_APPROACH_SELECTED,
			inputState:      MissionState_MISSION_STATE_APPROACH,
			resultState:     MissionState_MISSION_STATE_APPROACH,
			err:             ErrInvalidStateTransition,
		},
	}

	for _, scenario := range testCases {
		t.Run(scenario.desc, func(t *testing.T) {
			result, err := ResolveState(scenario.stateTransition, scenario.inputState)
			assert.IsType(t, scenario.err, err)
			assert.Equal(t, scenario.resultState, result)
		})
	}
}

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
			desc:            "transition from approach state to following state with target_approached",
			stateTransition: StateTransition_STATE_TRANSITION_TARGET_APPROACHED,
			inputState:      MissionState_MISSION_STATE_APPROACH,
			resultState:     MissionState_MISSION_STATE_FOLLOWING,
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
			desc:            "transition from following state to approach state with lock_failed",
			stateTransition: StateTransition_STATE_TRANSITION_LOCK_FAILED,
			inputState:      MissionState_MISSION_STATE_FOLLOWING,
			resultState:     MissionState_MISSION_STATE_APPROACH,
			err:             nil,
		},
		{
			desc:            "transition from following state to approach state with lock_success",
			stateTransition: StateTransition_STATE_TRANSITION_LOCK_SUCCESS,
			inputState:      MissionState_MISSION_STATE_FOLLOWING,
			resultState:     MissionState_MISSION_STATE_APPROACH,
			err:             nil,
		},
		{
			desc:            "transition from following state to kamikaze state with mode kamikaze selected",
			stateTransition: StateTransition_STATE_TRANSITION_MODE_KAMIKAZE_SELECTED,
			inputState:      MissionState_MISSION_STATE_FOLLOWING,
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
