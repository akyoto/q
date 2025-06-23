package core

import (
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/token"
)

// Compile turns a function into machine code.
func (f *Function) Compile() {
	registerCount := 0

	for _, input := range f.Input {
		f.Identifiers[input.Name] = f.AppendRegister(cpu.Register(registerCount))
		registerCount++

		if input.TypeTokens[0].Kind == token.ArrayStart {
			f.Identifiers[input.Name+".length"] = f.AppendRegister(cpu.Register(registerCount))
			registerCount++
		}
	}

	for instr := range f.Body.Instructions {
		f.Err = f.CompileInstruction(instr)

		if f.Err != nil {
			return
		}
	}

	f.Err = f.CheckDeadCode()
}