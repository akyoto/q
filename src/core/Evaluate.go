package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// Evaluate converts an expression to an SSA value.
func (f *Function) Evaluate(expr *expression.Expression) (*ssa.Value, error) {
	if expr.IsLeaf() {
		switch expr.Token.Kind {
		case token.Identifier:
			name := expr.Token.String(f.File.Bytes)
			value, exists := f.Identifiers[name]

			if !exists {
				return nil, errors.New(&UnknownIdentifier{Name: name}, f.File, expr.Token.Position)
			}

			return value, nil

		case token.Number:
			number, err := f.ToNumber(expr.Token)

			if err != nil {
				return nil, err
			}

			return f.AppendInt(number), nil

		case token.String:
			data := expr.Token.Bytes(f.File.Bytes)
			data = Unescape(data)
			return f.AppendBytes(data), nil
		}

		return nil, errors.New(InvalidExpression, f.File, expr.Token.Position)
	}

	switch expr.Token.Kind {
	case token.Call:
		children := expr.Children
		typ := ssa.Call

		if children[0].Token.Kind == token.Identifier {
			funcName := children[0].String(f.File.Bytes)

			if funcName == "len" {
				identifier := children[1].String(f.File.Bytes)
				return f.Identifiers[identifier+".length"], nil
			}

			if funcName == "syscall" {
				children = children[1:]
				typ = ssa.Syscall
			}
		}

		args := make([]*ssa.Value, len(children))

		for i, child := range children {
			value, err := f.Evaluate(child)

			if err != nil {
				return nil, err
			}

			args[i] = value
		}

		call := f.Append(ssa.Value{Type: typ, Args: args})
		return call, nil

	case token.Dot:
		name := fmt.Sprintf("%s.%s", expr.Children[0].String(f.File.Bytes), expr.Children[1].String(f.File.Bytes))
		return f.AppendFunction(name), nil
	}

	return nil, nil
}