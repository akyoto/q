package ssa_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

func TestStruct(t *testing.T) {
	fn := ssa.IR{}
	fn.AddBlock(ssa.NewBlock(""))
	hello := []byte("Hello")
	pointer := fn.Append(&ssa.Bytes{Bytes: hello})
	length := fn.Append(&ssa.Int{Int: len(hello)})
	fn.Append(&ssa.Struct{Typ: types.String, Arguments: ssa.Arguments{pointer, length}})
}