package types

// Slice creates a new slice type which is just a struct with a pointer and a length.
func Slice(typ Type, name string) *Struct {
	return &Struct{
		Package:    "",
		UniqueName: name,
		name:       name,
		Fields: []*Field{
			{Name: "ptr", Type: &Pointer{To: typ}, Index: 0, Offset: 0},
			{Name: "len", Type: UInt, Index: 1, Offset: 8},
		},
	}
}