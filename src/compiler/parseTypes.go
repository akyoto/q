package compiler

import (
	"iter"
	"strings"

	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// parseTypes parses the tokens of the input and output types.
func parseTypes(functions iter.Seq[*core.Function], env *core.Environment) error {
	for f := range functions {
		f.Type = &types.Function{
			Input:  make([]types.Type, len(f.Input)),
			Output: make([]types.Type, len(f.Output)),
		}

		for i, input := range f.Input {
			input.Typ = core.ParseType(input.Tokens[1:], f.File.Bytes, env)

			if input.Typ == nil {
				return errors.New(&UnknownType{Name: input.Tokens[1:].String(f.File.Bytes)}, f.File, input.Tokens[1].Position)
			}

			f.Type.Input[i] = input.Typ
		}

		for i, output := range f.Output {
			typeTokens := output.Tokens

			if len(output.Tokens) > 1 && output.Tokens[0].Kind == token.Identifier {
				output.Name = output.Tokens[0].String(f.File.Bytes)
				typeTokens = typeTokens[1:]
			}

			output.Typ = core.ParseType(typeTokens, f.File.Bytes, env)

			if output.Typ == nil {
				return errors.New(&UnknownType{Name: typeTokens.String(f.File.Bytes)}, f.File, typeTokens[0].Position)
			}

			f.Type.Output[i] = output.Typ
		}

		if f.Previous != nil || f.Next != nil {
			suffix := strings.Builder{}
			suffix.WriteByte('[')

			for i, input := range f.Input {
				suffix.WriteString(input.Typ.Name())

				if i != len(f.Input)-1 {
					suffix.WriteByte(',')
				}
			}

			suffix.WriteByte(']')
			f.Name += suffix.String()
			f.FullName += suffix.String()
		}
	}

	return nil
}