package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/fold"
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

	var folded map[ssa.Value]struct{}

	if f.Env.Build.FoldConstants {
		folded = fold.Constants(f.IR)
	}

	f.Finalize()
	err = f.removeDeadCode(folded)

	if err != nil {
		f.Err = err
		return
	}

	f.Err = f.checkResources()
}