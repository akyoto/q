package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestStruct(t *testing.T) {
	fn := ssa.IR{}
	hello := []byte("Hello")
	pointer := fn.Append(&ssa.Bytes{Bytes: hello})
	length := fn.Append(&ssa.Int{Int: len(hello)})
	str := fn.Append(&ssa.Struct{Typ: types.String, Arguments: ssa.Arguments{pointer, length}})
	assert.Equal(t, str.String(), "string{\"Hello\", 5}")
}