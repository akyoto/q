package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/set"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/ssa2asm"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// Function is the smallest unit of code.
type Function struct {
	ssa.IR
	ssa2asm.Compiler
	Name         string
	Package      string
	File         *fs.File
	Input        []*ssa.Parameter
	Output       []*ssa.Parameter
	Body         token.List
	Identifiers  map[string]ssa.Value
	All          *Environment
	Dependencies set.Ordered[*Function]
	Type         *types.Function
	Err          error
}

// NewFunction creates a new function.
func NewFunction(name string, pkg string, file *fs.File) *Function {
	return &Function{
		Name:        name,
		Package:     pkg,
		File:        file,
		Identifiers: make(map[string]ssa.Value, 8),
		IR: ssa.IR{
			Blocks: []*ssa.Block{
				{Instructions: make([]ssa.Value, 0, 8)},
			},
		},
		Compiler: ssa2asm.Compiler{
			UniqueName: fmt.Sprintf("%s.%s", pkg, name),
			Assembler: asm.Assembler{
				Instructions: make([]asm.Instruction, 0, 8),
			},
		},
	}
}

// EachDependency recursively finds all the calls to other functions.
// It avoids calling the same function twice with the help of a hashmap.
func (f *Function) EachDependency(traversed map[*Function]bool, call func(*Function)) {
	call(f)
	traversed[f] = true

	for dep := range f.Dependencies.All() {
		if traversed[dep] {
			continue
		}

		dep.EachDependency(traversed, call)
	}
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
	return f.UniqueName
}