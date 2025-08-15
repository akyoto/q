package types

import "git.urbach.dev/cli/q/src/fs"

// Struct is a structure in memory whose regions are addressable with named fields.
type Struct struct {
	File       *fs.File
	Package    string
	UniqueName string
	name       string
	Fields     []*Field
}

// NewStruct creates a new struct.
func NewStruct(file *fs.File, pkg string, name string) *Struct {
	return &Struct{
		File:       file,
		Package:    pkg,
		UniqueName: pkg + "." + name,
		name:       name,
	}
}

// AddField adds a new field to the end of the struct.
func (s *Struct) AddField(field *Field) {
	s.Fields = append(s.Fields, field)
}

// FieldByName returns the field with the given name if it exists.
func (s *Struct) FieldByName(name string) *Field {
	for _, field := range s.Fields {
		if field.Name == name {
			return field
		}
	}

	return nil
}

// Name returns the name of the struct.
func (s *Struct) Name() string {
	return s.name
}

// Size returns the total size in bytes.
func (s *Struct) Size() int {
	sum := 0

	for _, field := range s.Fields {
		sum += field.Type.Size()
	}

	return sum
}