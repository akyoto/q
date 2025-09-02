package types

import "strings"

// Tuple is a collection of multiple types.
type Tuple struct {
	Types []Type
}

// Name returns the type name.
func (t *Tuple) Name() string {
	b := strings.Builder{}
	b.WriteString("(")

	for i, typ := range t.Types {
		b.WriteString(typ.Name())

		if i != len(t.Types)-1 {
			b.WriteString(", ")
		}
	}

	b.WriteString(")")
	return b.String()
}

// Size returns the total size in bytes.
func (t *Tuple) Size() int {
	sum := 0

	for _, typ := range t.Types {
		sum += typ.Size()
	}

	return sum
}