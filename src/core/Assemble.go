package core

// Assemble translates SSA form into machine code.
func (f *Function) Assemble() {
	f.GenerateAssembly(f.IR, f.Env.Build, f.needsStackFrame(), f.Assembler.Libraries.Count() > 0)
}