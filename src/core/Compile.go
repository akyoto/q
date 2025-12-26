package core

import (
	"git.urbach.dev/cli/q/src/ast"
)

// Compile translates tokens to SSA form.
func (f *Function) Compile() {
	f.compileInputs()

	// From the body tokens we generate the AST which is
	// a list of top-level instructions.
	tree, err := ast.Parse(f.Body(), f.File)

	if err != nil {
		f.Err = err
		return
	}

	// Compile the AST nodes to SSA form.
	// This evaluates expressions and adds their SSA values
	// to basic blocks in the intermediate representation.
	err = f.compileAST(tree)

	if err != nil {
		f.Err = err
		return
	}

	f.Err = f.optimize()
}