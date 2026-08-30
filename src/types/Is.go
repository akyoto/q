package types

// Is returns true if the encountered type `a` can be converted into the expected type `b`.
func Is(a Type, b Type) bool {
	if a == b || b == Any {
		return true
	}

	aPointer, aIsPointer := a.(*Pointer)
	bPointer, bIsPointer := b.(*Pointer)

	if aIsPointer && bIsPointer && (aPointer.To == bPointer.To || bPointer.To == Any) {
		return true
	}

	bUnion, bIsUnion := b.(*Union)

	if bIsUnion {
		return bUnion.Index(a) != -1
	}

	aResource, aIsResource := a.(*Resource)

	if aIsResource {
		bResource, bIsResource := b.(*Resource)

		if bIsResource {
			return Is(aResource.Of, bResource.Of)
		}

		if Is(aResource.Of, b) {
			return true
		}
	}

	if a == AnyInt || a == Error {
		return isIntCompatible(b)
	}

	if b == AnyInt || b == Error {
		return isIntCompatible(a)
	}

	return false
}