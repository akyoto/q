package types

// Resource represents a shared resource that must be
// constructed, used and deconstructed, in that order.
type Resource struct {
	Of Type
}

// Name returns the type name.
func (r *Resource) Name() string {
	return "!" + r.Of.Name()
}

// Size returns the total size in bytes.
func (r *Resource) Size() int {
	return r.Of.Size()
}