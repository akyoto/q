package compiler

import (
	"iter"

	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// parseTypes parses the tokens of the input and output types.
func parseTypes(functions iter.Seq[*core.Function], env *core.Environment) {
	for f := range functions {
		f.Type = &types.Function{
			Input:  make([]types.Type, len(f.Input)),
			Output: make([]types.Type, len(f.Output)),
		}

		for i, input := range f.Input {
			input.Typ = core.ParseType(input.Tokens[1:], f.File.Bytes, env)
			f.Type.Input[i] = input.Typ
		}

		for i, output := range f.Output {
			if len(output.Tokens) > 1 && output.Tokens[0].Kind == token.Identifier {
				output.Typ = core.ParseType(output.Tokens[1:], f.File.Bytes, env)
			} else {
				output.Typ = core.ParseType(output.Tokens, f.File.Bytes, env)
			}

			f.Type.Output[i] = output.Typ
		}
	}
}