package core

// functionOverloads contains a linked list of functions.
type functionOverloads struct {
	Previous *Function
	Next     *Function
}

// AddSuffix adds a suffix to the name and is used for generic functions.
func (f *Function) AddSuffix(suffix string) {
	f.name += suffix
	f.FullName += suffix
}

// Variants returns all function overloads.
func (f *Function) Variants(yield func(*Function) bool) {
	for {
		if !yield(f) {
			return
		}

		f = f.Next

		if f == nil {
			return
		}
	}
}