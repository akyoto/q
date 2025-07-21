package asm

import (
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/token"
)

type Instruction interface{}

type AddRegisterRegister struct {
	Destination cpu.Register
	Source      cpu.Register
	Operand     cpu.Register
}

type AndRegisterNumber struct {
	Destination cpu.Register
	Source      cpu.Register
	Number      int
}

type AndRegisterRegister struct {
	Destination cpu.Register
	Source      cpu.Register
	Operand     cpu.Register
}

type Call struct {
	Label string
}

type CallExtern struct {
	Library  string
	Function string
}

type CallExternStart struct{}
type CallExternEnd struct{}

type CompareRegisterRegister struct {
	SourceA cpu.Register
	SourceB cpu.Register
}

type DivRegisterRegister struct {
	Destination cpu.Register
	Source      cpu.Register
	Operand     cpu.Register
}

type Jump struct {
	Label     string
	Condition token.Kind
}

type Label struct {
	Name string
}

type ModRegisterRegister struct {
	Destination cpu.Register
	Source      cpu.Register
	Operand     cpu.Register
}

type MoveRegisterLabel struct {
	Label       string
	Destination cpu.Register
}

type MoveRegisterNumber struct {
	Destination cpu.Register
	Number      int
}

type MoveRegisterRegister struct {
	Destination cpu.Register
	Source      cpu.Register
}

type MulRegisterRegister struct {
	Destination cpu.Register
	Source      cpu.Register
	Operand     cpu.Register
}

type NegateRegister struct {
	Destination cpu.Register
	Source      cpu.Register
}

type OrRegisterRegister struct {
	Destination cpu.Register
	Source      cpu.Register
	Operand     cpu.Register
}

type PopRegisters struct {
	Registers []cpu.Register
}

type PushRegisters struct {
	Registers []cpu.Register
}

type Return struct{}

type SubRegisterRegister struct {
	Destination cpu.Register
	Source      cpu.Register
	Operand     cpu.Register
}

type SubRegisterNumber struct {
	Destination cpu.Register
	Source      cpu.Register
	Number      int
}

type StackFrameStart struct {
	FramePointer bool
	ExternCalls  bool
}

type StackFrameEnd struct {
	FramePointer bool
}

type Syscall struct{}

type XorRegisterRegister struct {
	Destination cpu.Register
	Source      cpu.Register
	Operand     cpu.Register
}