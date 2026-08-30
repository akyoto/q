package core

import (
	"strings"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// MaxInputRegisters is the portable register limit.
const MaxInputRegisters = 7

// compileInputs registers every input as an identifier.
func (f *Function) compileInputs() error {
	offset := 0
	registers := 0

	for i, input := range f.Input {
		structType, isStructType := types.Unwrap(input.Typ).(*types.Struct)
		registers += inputRegisterCount(input)

		if strings.HasPrefix(input.Name, "_") {
			if isStructType {
				offset += len(structType.Fields) - 1
			}

			continue
		}

		if isStructType {
			var structure *ssa.Struct
			structure, offset = f.composeStruct(structType, input, i, offset)
			f.Block().Identify(input.Name, structure)
			continue
		}

		input.Index = uint8(offset + i)
		f.Block().Identify(input.Name, input)
		f.Append(input)
	}

	if registers > MaxInputRegisters {
		return errors.New(&RegisterLimitExceeded{Function: f.Name(), Required: registers, Available: MaxInputRegisters}, f.File, f.Input[len(f.Input)-1].Source)
	}

	return nil
}
