package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// TypeFromTokens returns the type with the given tokens or `nil` if it doesn't exist.
func TypeFromTokens(tokens token.List, file *fs.File, env *Environment) (types.Type, error) {
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

		typ, err := TypeFromTokens(tokens[:i], file, env)

		if err != nil {
			return nil, err
		}

		union.Types = append(union.Types, typ)
		tokens = tokens[i+1:]
	}

	if union != nil {
		typ, err := TypeFromTokens(tokens, file, env)

		if err != nil {
			return nil, err
		}

		union.Types = append(union.Types, typ)
		return union, nil
	}

	if tokens[0].Kind == token.Not {
		typ, err := TypeFromTokens(tokens[1:], file, env)

		if err != nil {
			return nil, err
		}

		return &types.Resource{Of: typ}, nil
	}

	if tokens[0].Kind == token.Mul {
		typ, err := TypeFromTokens(tokens[1:], file, env)

		if err != nil {
			return nil, err
		}

		if typ == types.Any {
			return types.AnyPointer, nil
		}

		return &types.Pointer{To: typ}, nil
	}

	if len(tokens) >= 2 && tokens[0].Kind == token.ArrayStart && tokens[1].Kind == token.ArrayEnd {
		return nil, errors.New(&NotImplemented{Subject: "array types"}, file, tokens[0].Position)
	}

	if tokens[0].Kind != token.Identifier {
		return nil, errors.New(&UnknownType{Name: tokens.String(file.Bytes)}, file, tokens[0].Position)
	}

	name := tokens[0].String(file.Bytes)

	switch name {
	case "string":
		return types.String, nil
	case "int":
		return types.Int, nil
	case "int64":
		return types.Int64, nil
	case "int32":
		return types.Int32, nil
	case "int16":
		return types.Int16, nil
	case "int8":
		return types.Int8, nil
	case "uint":
		return types.UInt, nil
	case "uint64":
		return types.UInt64, nil
	case "uint32":
		return types.UInt32, nil
	case "uint16":
		return types.UInt16, nil
	case "uint8":
		return types.UInt8, nil
	case "byte":
		return types.Byte, nil
	case "bool":
		return types.Bool, nil
	case "float":
		return types.Float, nil
	case "float64":
		return types.Float64, nil
	case "float32":
		return types.Float32, nil
	case "error":
		return types.Error, nil
	case "nil":
		return types.Nil, nil
	case "any":
		return types.Any, nil
	}

	if env != nil {
		// TODO: Optimize this and check for the correct package.
		for structure := range env.Structs() {
			if structure.Name() == name {
				return structure, nil
			}
		}
	}

	return nil, errors.New(&UnknownType{Name: tokens.String(file.Bytes)}, file, tokens[0].Position)
}