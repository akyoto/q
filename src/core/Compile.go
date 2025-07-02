package core

import "git.urbach.dev/cli/q/src/types"

// Compile turns a function into machine code.
func (f *Function) Compile() {
	offset := 0

	for i, input := range f.Input {
		if input.Name == "_" {
			continue
		}

		input.Index = uint8(offset + i)
		f.Append(input)
		f.Identifiers[input.Name] = input
		structure, isStruct := input.Typ.(*types.Struct)

		if isStruct {
			offset += len(structure.Fields) - 1
		}
	}

	for instr := range f.Body.Instructions {
		f.Err = f.CompileInstruction(instr)

		if f.Err != nil {
			return
		}
	}

	f.Err = f.CheckDeadCode()

	if f.Err != nil {
		return
	}

	f.GenerateAssembly(f.IR, f.IsLeaf())
}