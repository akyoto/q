package types

// isIntCompatible returns true if the type is integer compatible.
func isIntCompatible(t Type) bool {
	switch t {
	case Int64, Int32, Int16, Int8, UInt64, UInt32, UInt16, UInt8, Error, Nil, AnyInt:
		return true
	default:
		return false
	}
}