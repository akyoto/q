package core

import (
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// ParseType returns the type with the given tokens or `nil` if it doesn't exist.
func ParseType[T ~[]token.Token](tokens T, source []byte, env *Environment) types.Type {
	if len(tokens) == 0 {
		return nil
	}

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

		union.Types = append(union.Types, ParseType(tokens[:i], source, env))
		tokens = tokens[i+1:]
	}

	if union != nil {
		union.Types = append(union.Types, ParseType(tokens, source, env))
		return union
	}

	if tokens[0].Kind == token.Not {
		to := tokens[1:]
		typ := ParseType(to, source, env)

		if typ == nil {
			return nil
		}

		return &types.Resource{Of: typ}
	}

	if tokens[0].Kind == token.Mul {
		to := tokens[1:]
		typ := ParseType(to, source, env)

		if typ == nil {
			return nil
		}

		if typ == types.Any {
			return types.AnyPointer
		}

		return &types.Pointer{To: typ}
	}

	if len(tokens) >= 2 && tokens[0].Kind == token.ArrayStart && tokens[1].Kind == token.ArrayEnd {
		to := tokens[2:]
		typ := ParseType(to, source, env)

		if typ == nil {
			return nil
		}

		if typ == types.Any {
			return types.AnyArray
		}

		return &types.Array{Of: typ}
	}

	if tokens[0].Kind != token.Identifier {
		return nil
	}

	name := tokens[0].String(source)

	switch name {
	case "string":
		return types.String
	case "int":
		return types.Int
	case "int64":
		return types.Int64
	case "int32":
		return types.Int32
	case "int16":
		return types.Int16
	case "int8":
		return types.Int8
	case "uint":
		return types.UInt
	case "uint64":
		return types.UInt64
	case "uint32":
		return types.UInt32
	case "uint16":
		return types.UInt16
	case "uint8":
		return types.UInt8
	case "byte":
		return types.Byte
	case "bool":
		return types.Bool
	case "float":
		return types.Float
	case "float64":
		return types.Float64
	case "float32":
		return types.Float32
	case "any":
		return types.Any
	}

	if env == nil {
		return nil
	}

	// TODO: Optimize this and check for the correct package.
	for structure := range env.Structs() {
		if structure.Name() == name {
			return structure
		}
	}

	return nil
}