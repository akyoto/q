package types

// Is returns true if the encountered type `a` can be converted into the expected type `b`.
func Is(a Type, b Type) bool {
	if a == b || b == Any || a == nil {
		return true
	}

	aPointer, aIsPointer := a.(*Pointer)
	bPointer, bIsPointer := b.(*Pointer)

	if aIsPointer && bIsPointer && (bPointer.To == Any || aPointer.To == bPointer.To) {
		return true
	}

	aArray, aIsArray := a.(*Array)

	if aIsArray && bIsPointer && (bPointer.To == Any || aArray.Of == bPointer.To) {
		return true
	}

	bArray, bIsArray := b.(*Array)

	if aIsArray && bIsArray && (bArray.Of == Any || aArray.Of == bArray.Of) {
		return true
	}

	if a == AnyInt {
		switch b {
		case Int64, Int32, Int16, Int8, UInt64, UInt32, UInt16, UInt8, AnyInt:
			return true
		default:
			return false
		}
	}

	if b == AnyInt {
		switch a {
		case Int64, Int32, Int16, Int8, UInt64, UInt32, UInt16, UInt8, AnyInt:
			return true
		default:
			return false
		}
	}

	return false
}