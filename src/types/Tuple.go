package types

import "fmt"

// Tuple is a collection of multiple types.
type Tuple struct {
	Types []Type
}

// Name returns the type name.
func (t *Tuple) Name() string {
	return fmt.Sprintf("[%d]tuple", len(t.Types))
}

// Size returns the total size in bytes.
func (t *Tuple) Size() int {
	sum := 0

	for _, typ := range t.Types {
		sum += typ.Size()
	}

	return sum
}