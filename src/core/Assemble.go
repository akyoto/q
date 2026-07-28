package core

// Assemble translates SSA form into machine code.
func (f *Function) Assemble() {
	for global := range f.Globals.All() {
		global.Used.Add(1)
	}

	f.CompileToAssembly(f.IR, f.Env.Build, f.needsStackFrame(), f.Assembler.Libraries.Count() > 0)
}