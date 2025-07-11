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

type PushRegister struct {
	Register cpu.Register
}

type Return struct{}

type SubRegisterNumber struct {
	Destination cpu.Register
	Source      cpu.Register
	Number      int
}

type StackFrameStart struct{}
type StackFrameEnd struct{}

type Syscall struct{}