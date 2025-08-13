package codegen

// isLastStepInBlock returns true if the given `step` is the last step in the block.
func (f *Function) isLastStepInBlock(step *Step) bool {
	return step.Index == len(f.Steps)-1 || step.Block != f.Steps[step.Index+1].Block
}