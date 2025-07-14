package asm

import "git.urbach.dev/cli/q/src/cpu"

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

type Call struct {
	Label string
}

type CallExtern struct {
	Library  string
	Function string
}

type CallExternStart struct{}
type CallExternEnd struct{}

type DivRegisterRegister struct {
	Destination cpu.Register
	Source      cpu.Register
	Operand     cpu.Register
}

type Jump struct {
	Label string
}

type Label struct {
	Name string
}

type MoveRegisterLabel struct {
	Destination cpu.Register
	Label       string
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