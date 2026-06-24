package core

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateSyscall converts a syscall to an SSA value.
func (f *Function) evaluateSyscall(expr *expression.Expression) (ssa.Value, error) {
	if f.Env.Build.OS == config.Windows {
		return nil, errors.New(SyscallNotAvailable, f.File, expr.Children[0].Source())
	}

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