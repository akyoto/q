package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// evaluateCallPointer converts a function pointer call to an SSA value.
func (f *Function) evaluateCallPointer(funcValue ssa.Value, source token.Source) (ssa.Value, error) {
	call := f.Append(&ssa.CallPointer{
		Arguments: ssa.Arguments{funcValue},
		Source:    source,
	})

	return call, nil
}