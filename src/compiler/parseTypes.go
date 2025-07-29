package compiler

import (
	"iter"

	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/types"
)

// parseTypes parses the tokens of the input and output types.
func parseTypes(functions iter.Seq[*core.Function]) {
	for f := range functions {
		f.Type = &types.Function{
			Input:  make([]types.Type, len(f.Input)),
			Output: make([]types.Type, len(f.Output)),
		}

		for i, input := range f.Input {
			input.Typ = types.Parse(input.Tokens[1:], f.File.Bytes)
			f.Type.Input[i] = input.Typ
		}

		for i, output := range f.Output {
			if len(output.Tokens) > 1 {
				output.Typ = types.Parse(output.Tokens[1:], f.File.Bytes)
			} else {
				output.Typ = types.Parse(output.Tokens, f.File.Bytes)
			}

			f.Type.Output[i] = output.Typ
		}
	}
}