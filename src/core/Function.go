package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// Function is the smallest unit of code.
type Function struct {
	name   string
	file   *fs.File
	Input  []*Parameter
	Output []*Parameter
	Body   token.List
}

// NewFunction creates a new function.
func NewFunction(name string, file *fs.File) *Function {
	return &Function{
		name: name,
		file: file,
	}
}

// IsExtern returns true if the function has no body.
func (f *Function) IsExtern() bool {
	return f.Body == nil
}

// ResolveTypes parses the input and output types.
func (f *Function) ResolveTypes() error {
	for _, param := range f.Input {
		param.name = param.tokens[0].String(f.file.Bytes)
	}

	return nil
}

// String returns the package and function name.
func (f *Function) String() string {
	return fmt.Sprintf("%s.%s", f.file.Package, f.name)
}