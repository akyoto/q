package core

// Compile turns a function into machine code.
func (f *Function) Compile() {
	f.registerInputs()
	err := f.compileTokens(f.Body)

	if err != nil {
		f.Err = err
		return
	}

	f.Finalize()
	err = f.checkDeadCode()

	if err != nil {
		f.Err = err
		return
	}

	f.GenerateAssembly(f.IR, f.needsStackFrame(), f.Assembler.Libraries.Count() > 0)
}