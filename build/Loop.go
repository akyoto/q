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
	state.loop.counter++
	label := fmt.Sprintf("loop_%d", state.loop.counter)
	state.loop.labels = append(state.loop.labels, label)
	state.assembler.AddLabel(label)
	return nil
}

// LoopEnd handles the end of loops.
func (state *State) LoopEnd() error {
	state.scopes.Pop()
	label := state.loop.labels[len(state.loop.labels)-1]
	state.assembler.Jump(label)
	state.loop.labels = state.loop.labels[:len(state.loop.labels)-1]
	return nil
}
