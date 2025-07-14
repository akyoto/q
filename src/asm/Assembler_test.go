package asm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestAssembler(t *testing.T) {
	a := &asm.Assembler{}
	a.Append(&asm.Label{Name: "a"})
	a.Append(&asm.StackFrameStart{})
	a.Append(&asm.Jump{Label: "b"})
	a.Append(&asm.Jump{Label: "c"})
	a.Append(&asm.Jump{Label: "d"})
	a.Append(&asm.Jump{Label: "e"})
	a.Append(&asm.Jump{Label: "f"})
	a.Append(&asm.Call{Label: "b"})
	a.Append(&asm.Call{Label: "c"})
	a.Append(&asm.Call{Label: "d"})
	a.Append(&asm.Call{Label: "e"})
	a.Append(&asm.Call{Label: "f"})
	a.Append(&asm.StackFrameEnd{})
	a.Append(&asm.Return{})

	b := &asm.Assembler{}
	b.Append(&asm.Label{Name: "b"})
	b.Append(&asm.AddRegisterRegister{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.SubRegisterRegister{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.MulRegisterRegister{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.DivRegisterRegister{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.AddRegisterRegister{Destination: 1, Source: 1, Operand: 2})
	b.Append(&asm.SubRegisterRegister{Destination: 1, Source: 1, Operand: 2})
	b.Append(&asm.MulRegisterRegister{Destination: 1, Source: 1, Operand: 2})
	b.Append(&asm.DivRegisterRegister{Destination: 1, Source: 1, Operand: 2})
	b.Append(&asm.Return{})

	c := &asm.Assembler{}
	c.Append(&asm.Label{Name: "c"})
	c.Append(&asm.PushRegisters{Registers: []cpu.Register{0}})
	c.Append(&asm.PushRegisters{Registers: []cpu.Register{0, 1}})
	c.Append(&asm.PushRegisters{Registers: []cpu.Register{0, 1, 2}})
	c.Append(&asm.PushRegisters{Registers: []cpu.Register{0, 1, 2, 3}})
	c.Append(&asm.PushRegisters{Registers: []cpu.Register{0, 1, 2, 3, 4}})
	c.Append(&asm.PopRegisters{Registers: []cpu.Register{0, 1, 2, 3, 4}})
	c.Append(&asm.PopRegisters{Registers: []cpu.Register{0, 1, 2, 3}})
	c.Append(&asm.PopRegisters{Registers: []cpu.Register{0, 1, 2}})
	c.Append(&asm.PopRegisters{Registers: []cpu.Register{0, 1}})
	c.Append(&asm.PopRegisters{Registers: []cpu.Register{0}})
	c.Append(&asm.Return{})

	d := &asm.Assembler{}
	d.Append(&asm.Label{Name: "d"})
	d.Append(&asm.MoveRegisterLabel{Destination: 0, Label: "a"})
	d.Append(&asm.MoveRegisterLabel{Destination: 0, Label: "b"})
	d.Append(&asm.MoveRegisterLabel{Destination: 0, Label: "c"})
	d.Append(&asm.MoveRegisterLabel{Destination: 0, Label: "d"})
	d.Append(&asm.MoveRegisterLabel{Destination: 0, Label: "e"})
	d.Append(&asm.MoveRegisterNumber{Destination: 0, Number: 123})
	d.Append(&asm.MoveRegisterRegister{Destination: 0, Source: 1})
	d.Append(&asm.Return{})

	e := &asm.Assembler{}
	e.Append(&asm.Label{Name: "e"})
	a.Append(&asm.StackFrameStart{})
	e.Append(&asm.Jump{Label: "branch2"})
	e.Append(&asm.Label{Name: "branch1"})
	e.Append(&asm.Label{Name: "branch2"})
	e.Append(&asm.Syscall{})
	e.Append(&asm.StackFrameEnd{})
	e.Append(&asm.Return{})

	f := &asm.Assembler{}
	f.Libraries.Append("kernel32", "ExitProcess")
	f.Append(&asm.Label{Name: "f"})
	f.Append(&asm.StackFrameStart{FramePointer: true, ExternCalls: true})
	f.Append(&asm.CallExtern{Library: "kernel32", Function: "ExitProcess"})
	f.Append(&asm.StackFrameEnd{FramePointer: true})
	f.Append(&asm.Return{})

	final := asm.Assembler{}
	final.Merge(a)
	final.Merge(b)
	final.Merge(c)
	final.Merge(d)
	final.Merge(e)
	final.Merge(f)

	code, _, _ := final.Compile(&config.Build{Arch: config.ARM})
	assert.NotNil(t, code)

	code, _, _ = final.Compile(&config.Build{Arch: config.X86})
	assert.NotNil(t, code)
}