package types

// IsUnsigned returns true if the type represents an unsigned value.
func IsUnsigned(typ Type) bool {
	_, isPointer := typ.(*Pointer)
	return isPointer || typ == UInt64 || typ == UInt32 || typ == UInt16 || typ == UInt8
}