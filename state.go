package mastermind

import (
	"errors"

	pb "github.com/koustech/mastermind/gen/proto/go/mastermind/v1"
)

// ErrInvalidStateTransition is the error returned by ResolveState when an invalid StateTransition happens
var ErrInvalidStateTransition = errors.New("current state has no defined behavior for requested state transition")

// ResolveState resolves the MissionState according to the current MissionState and StateTransition given as input
func ResolveState(state_transition pb.StateTransition, current_state pb.MissionState) (pb.MissionState, error) {
	// NOTE: could this be done as a map or similar structure? Would it be thread-safe?
	new_state := current_state
	return new_state, ErrInvalidStateTransition
}
