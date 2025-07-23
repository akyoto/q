package asm

import "git.urbach.dev/cli/q/src/cpu"

type (
	rr struct {
		Destination cpu.Register
		Source      cpu.Register
	}

	rrn struct {
		Destination cpu.Register
		Source      cpu.Register
		Number      int
	}

	rrr struct {
		Destination cpu.Register
		Source      cpu.Register
		Operand     cpu.Register
	}
)