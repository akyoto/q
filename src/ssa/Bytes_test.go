package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestBytes(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock(ssa.NewBlock("fn"))
	hello := fn.Append(&ssa.Bytes{Bytes: []byte("Hello")})
	world := fn.Append(&ssa.Bytes{Bytes: []byte("World")})
	helloDup := fn.Append(&ssa.Bytes{Bytes: []byte("Hello")})
	one := fn.Append(&ssa.Int{Int: 1})

	assert.False(t, hello.Equals(world))
	assert.False(t, hello.Equals(one))
	assert.True(t, hello.Equals(helloDup))
	assert.Equal(t, hello.String(), "\"Hello\"")
	assert.True(t, types.Is(hello.Type(), types.CString))
}