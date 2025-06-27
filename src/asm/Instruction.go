package asm

import "git.urbach.dev/cli/q/src/cpu"

type Instruction interface{}

type Call struct {
	Label string
}

type FunctionStart struct{}
type FunctionEnd struct{}

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

type Return struct{}
type Syscall struct{}