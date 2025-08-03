package codegen

// assignFreeRegister assigns a free registers to the step.
func (f *Function) assignFreeRegister(step *Step) {
	step.Register = f.findFreeRegister(step)
}