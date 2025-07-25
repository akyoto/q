package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestSyscall(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock(ssa.NewBlock(""))
	syscall := fn.Append(&ssa.Syscall{})
	one := fn.Append(&ssa.Int{Int: 1})
	syscall2 := fn.Append(&ssa.Syscall{Arguments: ssa.Arguments{one}})
	two := fn.Append(&ssa.Int{Int: 2})
	syscall3 := fn.Append(&ssa.Syscall{Arguments: ssa.Arguments{one, two}})
	syscall4 := fn.Append(&ssa.Syscall{Arguments: ssa.Arguments{one, two}})

	assert.True(t, syscall.Type() == types.Any)
	assert.False(t, syscall2.Equals(syscall))
	assert.False(t, syscall4.Equals(one))
	assert.False(t, syscall4.Equals(syscall))
	assert.True(t, syscall4.Equals(syscall3))
}