package asm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/go/assert"
)

func TestAssembler(t *testing.T) {
	a := &asm.Assembler{}
	a.Append(&asm.Label{Name: "a"})
	a.Append(&asm.StackFrameStart{})
	a.Append(&asm.Call{Label: "b"})
	a.Append(&asm.Call{Label: "c"})
	a.Append(&asm.MoveRegisterLabel{Label: "b"})
	a.Append(&asm.MoveRegisterNumber{Destination: 0, Number: 123})
	a.Append(&asm.MoveRegisterRegister{Destination: 0, Source: 1})
	a.Append(&asm.StackFrameEnd{})
	a.Append(&asm.Return{})

	b := &asm.Assembler{}
	b.Append(&asm.Label{Name: "b"})
	b.Append(&asm.Syscall{})
	b.Append(&asm.Return{})

	c := &asm.Assembler{}
	c.Append(&asm.Label{Name: "c"})
	c.Append(&asm.Jump{Label: "branch2"})
	c.Append(&asm.Label{Name: "branch1"})
	c.Append(&asm.Label{Name: "branch2"})
	c.Append(&asm.Return{})

	final := asm.Assembler{}
	final.Merge(a)
	final.Merge(b)
	final.Merge(c)

	code, _, _ := final.Compile(&config.Build{Arch: config.ARM})
	assert.NotNil(t, code)

	code, _, _ = final.Compile(&config.Build{Arch: config.X86})
	assert.NotNil(t, code)
}