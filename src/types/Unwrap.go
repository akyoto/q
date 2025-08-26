package types

// Unwrap returns the underlying type without a resource wrapper.
func Unwrap(wrapped Type) Type {
	resource, isResource := wrapped.(*Resource)

	if isResource {
		return resource.Of
	}

	return wrapped
}