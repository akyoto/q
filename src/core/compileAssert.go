package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// compileAssert compiles an assertion.
func (f *Function) compileAssert(assert *ast.Assert) error {
	cond, err := f.evaluate(assert.Condition)

	if err != nil {
		return err
	}

	comparison := cond.(*ssa.BinaryOp)
	left := comparison.Left

	if left.Type() == types.Error {
		right := comparison.Right.(*ssa.Int)

		switch {
		case assert.Condition.Token.Kind == token.NotEqual && right.Int == 0:
			for _, protected := range f.Block().Protected[left] {
				f.Block().Unidentify(protected)
			}

			f.Block().Unidentify(left)
			f.Block().Unprotect(left)
		case assert.Condition.Token.Kind == token.Equal && right.Int == 0:
			f.Block().Unidentify(left)
			f.Block().Unprotect(left)
		}
	}

	f.Append(&ssa.Assert{Condition: cond})
	f.Dependencies.Add(f.Env.Function("run", "crash"))
	return nil
}