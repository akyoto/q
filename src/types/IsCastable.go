package types

// IsCastable returns true if the `a` type can be casted to the `b` type using an explicit cast.
func IsCastable(a Type, b Type) bool {
	if a == Any {
		return true
	}

	_, aIsPointer := a.(*Pointer)
	_, bIsPointer := b.(*Pointer)

	if aIsPointer && bIsPointer {
		return true
	}

	aIsInt := Is(a, AnyInt)
	bIsInt := Is(b, AnyInt)

	if (aIsInt || aIsPointer) && bIsInt {
		return true
	}

	_, aIsFunction := a.(*Function)

	if aIsFunction && (bIsInt || bIsPointer) {
		return true
	}

	return false
}