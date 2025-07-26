package core

import (
	"git.urbach.dev/cli/q/src/ast"
)

// compileCall compiles a call instruction.
func (f *Function) compileCall(call *ast.Call) error {
	_, err := f.evaluate(call.Expression)
	return err
}