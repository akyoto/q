package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// Compile turns a function into machine code.
func (f *Function) Compile() {
	extra := 0

	for i, input := range f.Input {
		if input.Name == "_" {
			continue
		}

		array, isArray := input.Typ.(*types.Array)

		if isArray {
			pointer := &ssa.Parameter{
				Index:  uint8(i + extra),
				Name:   input.Name,
				Typ:    &types.Pointer{To: array.Of},
				Source: input.Source,
			}

			f.Append(pointer)
			f.Identifiers[pointer.Name] = pointer
			extra++

			length := &ssa.Parameter{
				Index:  uint8(i + extra),
				Name:   input.Name + ".len",
				Typ:    types.AnyInt,
				Source: input.Source,
			}

			f.Append(length)
			f.Identifiers[length.Name] = length
			continue
		}

		input.Index = uint8(i + extra)
		f.Append(input)
		f.Identifiers[input.Name] = input
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

	f.ssaToAsm()
}