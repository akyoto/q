package asm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/token"
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
	a.Append(&asm.Jump{Label: "g"})
	a.Append(&asm.Call{Label: "b"})
	a.Append(&asm.Call{Label: "c"})
	a.Append(&asm.Call{Label: "d"})
	a.Append(&asm.Call{Label: "e"})
	a.Append(&asm.Call{Label: "f"})
	a.Append(&asm.Call{Label: "g"})
	a.Append(&asm.StackFrameEnd{})
	a.Append(&asm.Return{})

	b := &asm.Assembler{}
	b.Append(&asm.Label{Name: "b"})
	b.Append(&asm.Add{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.Add{Destination: 1, Source: 1, Operand: 2})
	b.Append(&asm.AddNumber{Destination: 0, Source: 1, Number: 42})
	b.Append(&asm.AddNumber{Destination: 1, Source: 1, Number: 42})
	b.Append(&asm.And{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.And{Destination: 1, Source: 1, Operand: 2})
	b.Append(&asm.AndNumber{Destination: 0, Source: 1, Number: 42})
	b.Append(&asm.AndNumber{Destination: 1, Source: 1, Number: 42})
	b.Append(&asm.Divide{Destination: 0, Source: 1, Operand: 3})
	b.Append(&asm.Divide{Destination: 1, Source: 1, Operand: 3})
	b.Append(&asm.DivideSigned{Destination: 0, Source: 1, Operand: 3})
	b.Append(&asm.DivideSigned{Destination: 1, Source: 1, Operand: 3})
	b.Append(&asm.Modulo{Destination: 0, Source: 1, Operand: 3})
	b.Append(&asm.ModuloSigned{Destination: 0, Source: 1, Operand: 3})
	b.Append(&asm.Multiply{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.Multiply{Destination: 1, Source: 1, Operand: 2})
	b.Append(&asm.Or{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.Or{Destination: 1, Source: 1, Operand: 2})
	b.Append(&asm.OrNumber{Destination: 0, Source: 1, Number: 42})
	b.Append(&asm.OrNumber{Destination: 1, Source: 1, Number: 42})
	b.Append(&asm.ShiftLeft{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.ShiftLeftNumber{Destination: 0, Source: 1, Number: 8})
	b.Append(&asm.ShiftRight{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.ShiftRightNumber{Destination: 0, Source: 1, Number: 8})
	b.Append(&asm.ShiftRightSigned{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.ShiftRightSignedNumber{Destination: 0, Source: 1, Number: 8})
	b.Append(&asm.Subtract{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.Subtract{Destination: 1, Source: 1, Operand: 2})
	b.Append(&asm.SubtractNumber{Destination: 0, Source: 1, Number: 42})
	b.Append(&asm.Xor{Destination: 0, Source: 1, Operand: 2})
	b.Append(&asm.Xor{Destination: 1, Source: 1, Operand: 2})
	b.Append(&asm.XorNumber{Destination: 0, Source: 1, Number: 42})
	b.Append(&asm.XorNumber{Destination: 1, Source: 1, Number: 42})
	b.Append(&asm.Return{})

	c := &asm.Assembler{}
	c.Append(&asm.Label{Name: "c"})
	c.Append(&asm.Push{Registers: []cpu.Register{0}})
	c.Append(&asm.Push{Registers: []cpu.Register{0, 1}})
	c.Append(&asm.Push{Registers: []cpu.Register{0, 1, 2}})
	c.Append(&asm.Push{Registers: []cpu.Register{0, 1, 2, 3}})
	c.Append(&asm.Push{Registers: []cpu.Register{0, 1, 2, 3, 4}})
	c.Append(&asm.Pop{Registers: []cpu.Register{0, 1, 2, 3, 4}})
	c.Append(&asm.Pop{Registers: []cpu.Register{0, 1, 2, 3}})
	c.Append(&asm.Pop{Registers: []cpu.Register{0, 1, 2}})
	c.Append(&asm.Pop{Registers: []cpu.Register{0, 1}})
	c.Append(&asm.Pop{Registers: []cpu.Register{0}})
	c.Append(&asm.Return{})

	d := &asm.Assembler{}
	d.Append(&asm.Label{Name: "d"})
	d.Data.SetImmutable("message", []byte("Hello"))
	d.Append(&asm.MoveLabel{Destination: 0, Label: "a"})
	d.Append(&asm.MoveLabel{Destination: 0, Label: "b"})
	d.Append(&asm.MoveLabel{Destination: 0, Label: "c"})
	d.Append(&asm.MoveLabel{Destination: 0, Label: "d"})
	d.Append(&asm.MoveLabel{Destination: 0, Label: "e"})
	d.Append(&asm.MoveLabel{Destination: 0, Label: "message"})
	d.Append(&asm.MoveNumber{Destination: 0, Number: 123})
	d.Append(&asm.Move{Destination: 0, Source: 1})
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

	g := &asm.Assembler{}
	g.Append(&asm.Label{Name: "g"})
	g.Append(&asm.Call{Label: "a"})
	g.Append(&asm.Compare{Destination: 0, Source: 1})
	g.Append(&asm.Jump{Label: "a", Condition: token.Equal})
	g.Append(&asm.Jump{Label: "a", Condition: token.NotEqual})
	g.Append(&asm.Jump{Label: "a", Condition: token.Greater})
	g.Append(&asm.Jump{Label: "a", Condition: token.GreaterEqual})
	g.Append(&asm.Jump{Label: "a", Condition: token.Less})
	g.Append(&asm.Jump{Label: "a", Condition: token.LessEqual})
	g.Append(&asm.Return{})

	final := asm.Assembler{}
	final.Merge(a)
	final.Merge(b)
	final.Merge(c)
	final.Merge(d)
	final.Merge(e)
	final.Merge(f)
	final.Merge(g)

	code, _, _ := final.Compile(&config.Build{Arch: config.ARM})
	assert.NotNil(t, code)

	code, _, _ = final.Compile(&config.Build{Arch: config.X86})
	assert.NotNil(t, code)
}