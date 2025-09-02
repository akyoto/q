package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
)

// Compile translates tokens to SSA form.
func (f *Function) Compile() {
	f.compileInputs()
	tree, err := ast.Parse(f.Body, f.File)

	if err != nil {
		f.Err = err
		return
	}

	err = f.compileAST(tree)

	if err != nil {
		f.Err = err
		return
	}

	if f.needsReturn() {
		f.Block().Append(&ssa.Return{})
	}

	f.Err = f.optimize()
}