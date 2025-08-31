package core

import "git.urbach.dev/cli/q/src/ast"

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

	folded := foldConstants(f.IR)
	f.Finalize()
	err = f.removeDeadCode(folded)

	if err != nil {
		f.Err = err
		return
	}

	f.Err = f.checkResources()
}