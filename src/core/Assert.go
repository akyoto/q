package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// Assert compiles an assertion.
func (f *Function) Assert(tokens token.List) error {
	condExpr := expression.Parse(tokens[1:])
	cond, err := f.evaluate(condExpr)

	if err != nil {
		return err
	}

	f.Append(&ssa.Assert{Condition: cond})
	f.Dependencies.Add(f.All.Function("run", "crash"))
	return nil
}