package types

// IsSigned returns true if the type represents a signed value.
func IsSigned(typ Type) bool {
	return typ == Int64 || typ == Int32 || typ == Int16 || typ == Int8 || typ == Error || typ == AnyInt || typ == Any
}

// IsUnsigned returns true if the type represents an unsigned value.
func IsUnsigned(typ Type) bool {
	return !IsSigned(typ)
}