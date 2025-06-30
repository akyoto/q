package types

import (
	"git.urbach.dev/cli/q/src/token"
)

// Parse returns the type with the given tokens or `nil` if it doesn't exist.
func Parse[T ~[]token.Token](tokens T, source []byte) Type {
	if len(tokens) == 0 {
		return nil
	}

	if tokens[0].Kind == token.Mul {
		to := tokens[1:]
		typ := Parse(to, source)

		if typ == nil {
			return nil
		}

		if typ == Any {
			return AnyPointer
		}

		return &Pointer{To: typ}
	}

	if len(tokens) >= 2 && tokens[0].Kind == token.ArrayStart && tokens[1].Kind == token.ArrayEnd {
		to := tokens[2:]
		typ := Parse(to, source)

		if typ == nil {
			return nil
		}

		if typ == Any {
			return AnyArray
		}

		return &Array{Of: typ}
	}

	if tokens[0].Kind != token.Identifier {
		return nil
	}

	switch tokens[0].String(source) {
	case "int":
		return Int
	case "int64":
		return Int64
	case "int32":
		return Int32
	case "int16":
		return Int16
	case "int8":
		return Int8
	case "uint":
		return UInt
	case "uint64":
		return UInt64
	case "uint32":
		return UInt32
	case "uint16":
		return UInt16
	case "uint8":
		return UInt8
	case "byte":
		return Byte
	case "bool":
		return Bool
	case "float":
		return Float
	case "float64":
		return Float64
	case "float32":
		return Float32
	case "any":
		return Any
	default:
		return nil
	}
}