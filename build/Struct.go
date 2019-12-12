package build

import "github.com/akyoto/q/build/types"

// Field is a field in a data structure.
type Field Parameter

// Struct represents a data structure.
type Struct struct {
	types.Type
	Fields []*Field
}
