package types

// IsUnsigned returns true if the type represents an unsigned value.
func IsUnsigned(typ Type) bool {
	return typ == UInt64 || typ == UInt32 || typ == UInt16 || typ == UInt8
}