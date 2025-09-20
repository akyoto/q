package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// TypeFromTokens returns the type with the given tokens or `nil` if it doesn't exist.
func (env *Environment) TypeFromTokens(tokens token.List, file *fs.File) (types.Type, error) {
	var union *types.Union

	for i, t := range tokens {
		if t.Kind != token.Or {
			continue
		}

		if union == nil {
			union = &types.Union{
				Types: make([]types.Type, 0, 2),
			}
		}

		typ, err := env.TypeFromTokens(tokens[:i], file)

		if err != nil {
			return nil, err
		}

		union.Types = append(union.Types, typ)
		tokens = tokens[i+1:]
	}

	if union != nil {
		typ, err := env.TypeFromTokens(tokens, file)

		if err != nil {
			return nil, err
		}

		union.Types = append(union.Types, typ)
		return union, nil
	}

	if tokens[0].Kind == token.Not {
		typ, err := env.TypeFromTokens(tokens[1:], file)

		if err != nil {
			return nil, err
		}

		return env.Resource(typ), nil
	}

	if tokens[0].Kind == token.Mul {
		typ, err := env.TypeFromTokens(tokens[1:], file)

		if err != nil {
			return nil, err
		}

		return env.Pointer(typ), nil
	}

	if len(tokens) >= 2 && tokens[0].Kind == token.ArrayStart && tokens[1].Kind == token.ArrayEnd {
		typ, err := env.TypeFromTokens(tokens[2:], file)

		if err != nil {
			return nil, err
		}

		return env.Slice(typ), nil
	}

	if tokens[0].Kind != token.Identifier {
		return nil, errors.New(&UnknownType{Name: tokens.String(file.Bytes)}, file, tokens[0].Position)
	}

	name := tokens[0].String(file.Bytes)
	typ := TypeByName(name, env)

	if typ != nil {
		return typ, nil
	}

	return nil, errors.New(&UnknownType{Name: tokens.String(file.Bytes)}, file, tokens[0].Position)
}