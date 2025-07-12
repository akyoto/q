package core

// Compile turns a function into machine code.
func (f *Function) Compile() {
	f.registerInputs()

	for instr := range f.Body.Instructions {
		f.Err = f.compileInstruction(instr)

		if f.Err != nil {
			return
		}
	}

	f.Err = f.checkDeadCode()

	if f.Err != nil {
		return
	}

	f.GenerateAssembly(f.IR, f.needsStackFrame())
}