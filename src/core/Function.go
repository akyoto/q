package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// Function is the smallest unit of code.
type Function struct {
	Name       string
	UniqueName string
	File       *fs.File
	Input      []*Parameter
	Output     []*Parameter
	Body       token.List
}

// NewFunction creates a new function.
func NewFunction(name string, file *fs.File) *Function {
	return &Function{
		Name:       name,
		File:       file,
		UniqueName: fmt.Sprintf("%s.%s", file.Package, name),
	}
}

// IsExtern returns true if the function has no body.
func (f *Function) IsExtern() bool {
	return f.Body == nil
}

// String returns the unique name.
func (f *Function) String() string {
	return f.UniqueName
}