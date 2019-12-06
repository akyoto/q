package build

import "fmt"

// LoopState handles the state of loop compilation.
type LoopState struct {
	counter int
	labels  []string
}

// LoopStart handles the start of loops.
func (state *State) LoopStart() error {
	state.scopes.Push()
	state.loopState.counter++
	label := fmt.Sprintf("loop_%d", state.loopState.counter)
	state.loopState.labels = append(state.loopState.labels, label)
	state.assembler.AddLabel(label)
	return nil
}

// LoopEnd handles the end of loops.
func (state *State) LoopEnd() error {
	err := state.PopScope(true)

	if err != nil {
		return err
	}

	label := state.loopState.labels[len(state.loopState.labels)-1]
	state.assembler.Jump(label)
	state.loopState.labels = state.loopState.labels[:len(state.loopState.labels)-1]
	return nil
}
