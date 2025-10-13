package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/token"
)

// compileLoopControl compiles loop control statements.
func (f *Function) compileLoopControl(control *ast.LoopControl) error {
	call := control.Expression

	if call.Token.Kind != token.Call {
		return errors.New(InvalidExpression, f.File, control.Expression.Source())
	}

	if len(call.Children) != 1 {
		return errors.New(InvalidExpression, f.File, control.Expression.Source())
	}

	loop := f.loopStack.Current()

	switch call.Children[0].Token.StringFrom(f.File.Bytes) {
	case "next":
		f.loopNext(loop)
	case "restart":
		// reserved
	case "stop":
		f.jump(loop.Exit)
	default:
		panic("invalid")
	}

	return nil
}