package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileAssert compiles an assertion.
func (f *Function) compileAssert(assert *ast.Assert) error {
	cond, err := f.evaluate(assert.Condition)

	if err != nil {
		return err
	}

	f.Append(&ssa.Assert{Condition: cond})
	f.Dependencies.Add(f.Env.Function("run", "crash"))
	return nil
}