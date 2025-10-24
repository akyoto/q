package core

import (
	"git.urbach.dev/cli/q/src/types"
)

// typeByName returns the type with the given name or `nil` if it doesn't exist.
func typeByName(name string, env *Environment) types.Type {
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
	case "error":
		return types.Error
	case "nil":
		return types.Nil
	case "any":
		return types.Any
	}

	if env != nil {
		// TODO: Optimize this and check for the correct package.
		for structure := range env.Structs() {
			if structure.Name() == name {
				return structure
			}
		}
	}

	return nil
}