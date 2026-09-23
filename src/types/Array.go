package types

import (
	"fmt"
	"strconv"
)

// Array creates a new static array type.
func Array(element Type, count int) *Struct {
	name := fmt.Sprintf("[%d]%s", count, element.Name())
	fields := make([]*Field, count)
	elementSize := element.Size()

	for i := range fields {
		fields[i] = &Field{
			Type:   element,
			Name:   strconv.Itoa(i),
			Index:  uint64(i),
			Offset: uint64(i * elementSize),
		}
	}

	return &Struct{
		Package:    "",
		UniqueName: name,
		name:       name,
		Fields:     fields,
	}
}