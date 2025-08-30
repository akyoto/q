package core

import (
	"git.urbach.dev/cli/q/src/codegen"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/set"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// Function is the smallest unit of code.
type Function struct {
	Name         string
	Package      string
	File         *fs.File
	Type         *types.Function
	Err          error
	Env          *Environment
	Input        []*ssa.Parameter
	Output       []*ssa.Parameter
	Body         token.List
	Dependencies set.Ordered[*Function]
	Previous     *Function
	Next         *Function
	ssa.IR
	codegen.Function
}

// AddInput adds an input parameter.
func (f *Function) AddInput(tokens token.List, source token.Source) {
	f.Input = append(f.Input, &ssa.Parameter{
		Tokens: tokens,
		Source: source,
	})
}

// AddOutput adds an output parameter.
func (f *Function) AddOutput(tokens token.List, source token.Source) {
	f.Output = append(f.Output, &ssa.Parameter{
		Tokens: tokens,
		Source: source,
	})
}

// IsExtern returns true if the function has no body.
func (f *Function) IsExtern() bool {
	return f.Body == nil
}

// IsLeaf returns true if the function doesn't call other functions.
func (f *Function) IsLeaf() bool {
	return f.Dependencies.Count() == 0
}

// String returns the unique name.
func (f *Function) String() string {
	return f.FullName
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