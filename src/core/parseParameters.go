package core

import (
	"iter"
	"strings"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// parseParameters parses the tokens of the input and output types.
func (env *Environment) parseParameters(functions iter.Seq[*Function]) error {
	for f := range functions {
		f.Type = &types.Function{
			Input:  make([]types.Type, len(f.Input)),
			Output: make([]types.Type, len(f.Output)),
		}

		for i, input := range f.Input {
			input.Name = input.Tokens[0].StringFrom(f.File.Bytes)
			typ, err := env.TypeFromTokens(input.Tokens[1:], f.File)

			if err != nil {
				return err
			}

			input.Typ = typ
			f.Type.Input[i] = input.Typ
		}

		for i, output := range f.Output {
			typeTokens := output.Tokens

			if len(output.Tokens) > 1 && output.Tokens[0].Kind == token.Identifier && output.Tokens[1].Kind != token.Or {
				output.Name = output.Tokens[0].StringFrom(f.File.Bytes)
				output.SetEnd(output.Tokens[0].End())
				typeTokens = typeTokens[1:]
			}

			typ, err := env.TypeFromTokens(typeTokens, f.File)

			if err != nil {
				return err
			}

			output.Typ = typ
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
			f.AddSuffix(suffix.String())
		}
	}

	return nil
}