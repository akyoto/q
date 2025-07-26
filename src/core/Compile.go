package core

import "git.urbach.dev/cli/q/src/ast"

// Compile turns a function into machine code.
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

	f.Finalize()
	err = f.removeDeadCode()

	if err != nil {
		f.Err = err
		return
	}

	f.GenerateAssembly(f.IR, f.needsStackFrame(), f.Assembler.Libraries.Count() > 0)
}