package codegen

// assignFreeRegister assigns a free registers to the step.
func (f *Function) assignFreeRegister(step *step) {
	users := step.Value.Users()

	if len(users) == 0 {
		return
	}

	from := step.Index
	to := f.ValueToStep[users[len(users)-1]].Index

	if from > to {
		from, to = to, from
	}

	step.Register = f.findFreeRegister(f.Steps[from:to])
}