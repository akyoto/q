package asm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/go/assert"
)

func TestAssembler(t *testing.T) {
	a := &asm.Assembler{}
	a.Append(&asm.Label{Name: "a"})
	a.Append(&asm.Call{Label: "b"})
	a.Append(&asm.Return{})

	b := &asm.Assembler{}
	b.Append(&asm.Label{Name: "b"})
	b.Append(&asm.Syscall{})
	b.Append(&asm.Return{})

	final := asm.Assembler{}
	final.Merge(a)
	final.Merge(b)

	code, _ := final.Compile(&build.Build{Arch: build.ARM})
	assert.NotNil(t, code)

	code, _ = final.Compile(&build.Build{Arch: build.X86})
	assert.NotNil(t, code)
}