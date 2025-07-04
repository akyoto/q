package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// Compile turns a function into machine code.
func (f *Function) Compile() {
	offset := 0

	for i, input := range f.Input {
		if input.Name == "_" {
			continue
		}

		structType, isStructType := input.Typ.(*types.Struct)

		if isStructType {
			for _, field := range structType.Fields {
				param := &ssa.Parameter{
					Index:  uint8(offset + i),
					Name:   fmt.Sprintf("%s.%s", input.Name, field.Name),
					Typ:    field.Type,
					Source: input.Source,
				}

				f.Identifiers[param.Name] = param
				f.Append(param)
				offset++
			}

			offset--
			continue
		}

		input.Index = uint8(offset + i)
		f.Identifiers[input.Name] = input
		f.Append(input)
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