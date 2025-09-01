package types

// Is returns true if the encountered type `a` can be converted into the expected type `b`.
func Is(a Type, b Type) bool {
	if a == b || a == Any || b == Any {
		return true
	}

	bUnion, bIsUnion := b.(*Union)

	if bIsUnion {
		return bUnion.Index(a) != -1
	}

	aPointer, aIsPointer := a.(*Pointer)
	bPointer, bIsPointer := b.(*Pointer)

	if aIsPointer && bIsPointer && (bPointer.To == Any || aPointer.To == bPointer.To) {
		return true
	}

	// TODO: Remove this temporary hack to allow integers as pointers
	if b == AnyPointer && a == AnyInt {
		return true
	}

	aResource, aIsResource := a.(*Resource)

	if aIsResource && Is(b, aResource.Of) {
		return true
	}

	bResource, bIsResource := b.(*Resource)

	if aIsResource && bIsResource {
		return aResource.Of == bResource.Of
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

	// TODO: Remove temporary hack to allow int64 -> uint32 conversion
	if a == Int64 && (b == UInt64 || b == UInt32) {
		return true
	}

	// TODO: Remove temporary hack to allow uint32 -> int64 conversion
	if (a == UInt32 || a == UInt64) && b == Int64 {
		return true
	}

	return false
}