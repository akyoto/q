package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
)

// compileInstruction compiles a single instruction.
func (f *Function) compileInstruction(instr token.List) error {
	if instr[0].IsKeyword() {
		switch instr[0].Kind {
		case token.Return:
			return f.compileReturn(instr)
		}
	}

	expr := expression.Parse(instr)

	if expr.Token.Kind == token.Define {
		left := expr.Children[0]
		right := expr.Children[1]
		name := left.String(f.File.Bytes)
		value, err := f.eval(right)
		f.Identifiers[name] = value
		return err
	}

	_, err := f.eval(expr)
	return err
}