package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
)

// compileInstruction compiles a single instruction.
func (f *Function) compileInstruction(instr token.List) error {
	if instr[0].Kind.IsKeyword() {
		switch instr[0].Kind {
		case token.If:
			return f.If(instr)
		case token.Return:
			return f.Return(instr)
		default:
			panic("keyword not implemented")
		}
	}

	expr := expression.Parse(instr)

	switch expr.Token.Kind {
	case token.Define:
		return f.Definition(expr)
	case token.String:
		return errors.New(&UnusedValue{Value: expr.String(f.File.Bytes)}, f.File, expr.Token.Position)
	}

	_, err := f.eval(expr)
	return err
}