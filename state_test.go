package mastermind

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
		desc             string
		state_transition StateTransition
		input_state      MissionState
		result_state     MissionState
		err              error
	}{
		{
			desc:             "transition from approach state with approach_selected = ERR",
			state_transition: StateTransition_STATE_TRANSITION_MODE_APPROACH_SELECTED,
			input_state:      MissionState_MISSION_STATE_APPROACH,
			result_state:     MissionState_MISSION_STATE_APPROACH,
			err:              nil,
		},
	}

	for _, scenario := range testCases {
		t.Run(scenario.desc, func(t *testing.T) {
			result, err := ResolveState(scenario.state_transition, scenario.input_state)
			assert.IsType(t, scenario.err, err)
			assert.Equal(t, scenario.result_state, result)
		})
	}
}

// TestResolveStateValidTransitions checks that all the transitions are correct
func TestResolveStateValidTransitions(t *testing.T) {
	// Setup
	testCases := []struct {
		desc             string
		state_transition StateTransition
		input_state      MissionState
		result_state     MissionState
		err              error
	}{
		{
			desc:             "transition from approach state to following state with target_approached",
			state_transition: StateTransition_STATE_TRANSITION_TARGET_APPROACHED,
			input_state:      MissionState_MISSION_STATE_APPROACH,
			result_state:     MissionState_MISSION_STATE_FOLLOWING,
			err:              nil,
		},
		{
			desc:             "transition from following state to approach state with lock_failed",
			state_transition: StateTransition_STATE_TRANSITION_LOCK_FAILED,
			input_state:      MissionState_MISSION_STATE_FOLLOWING,
			result_state:     MissionState_MISSION_STATE_APPROACH,
			err:              nil,
		},
	}

	for _, scenario := range testCases {
		t.Run(scenario.desc, func(t *testing.T) {
			result, err := ResolveState(scenario.state_transition, scenario.input_state)
			require.NoError(t, err)
			assert.Equal(t, scenario.result_state, result)
		})
	}
}
