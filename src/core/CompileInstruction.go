package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
)

func (f *Function) CompileInstruction(instr token.List) error {
	if instr[0].IsKeyword() {
		switch instr[0].Kind {
		case token.Return:
			return f.CompileReturn(instr)
		}
	}

	expr := expression.Parse(instr)

	if expr == nil {
		return nil
	}

	if expr.Token.Kind == token.Define {
		name := expr.Children[0].String(f.File.Bytes)
		value, err := f.Evaluate(expr.Children[1])
		f.Identifiers[name] = value
		return err
	}

	_, err := f.Evaluate(expr)
	return err
}