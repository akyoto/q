package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateSyscall converts a syscall to an SSA value.
func (f *Function) evaluateSyscall(expr *expression.Expression) (ssa.Value, error) {
	args, err := f.decompose(expr.Children[1:], nil, false)

	if err != nil {
		return nil, err
	}

	syscall := &ssa.Syscall{
		Arguments: args,
		Source:    expr.Source(),
	}

	return f.Append(syscall), nil
}