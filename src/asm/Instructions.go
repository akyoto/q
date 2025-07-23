package asm

import (
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/token"
)

type (
	Add        rrr
	And        rrr
	AndNumber  rrn
	Call       struct{ Label string }
	CallExtern struct {
		Library  string
		Function string
	}
	CallExternStart struct{}
	CallExternEnd   struct{}
	Compare         rr
	Divide          rrr
	Jump            struct {
		Label     string
		Condition token.Kind
	}
	Label     struct{ Name string }
	Modulo    rrr
	Move      rr
	MoveLabel struct {
		Label       string
		Destination cpu.Register
	}
	MoveNumber struct {
		Destination cpu.Register
		Number      int
	}
	Multiply         rrr
	Negate           rr
	Or               rrr
	Pop              struct{ Registers []cpu.Register }
	Push             struct{ Registers []cpu.Register }
	Return           struct{}
	ShiftLeft        rrr
	ShiftRightSigned rrr
	Subtract         rrr
	SubtractNumber   rrn
	StackFrameStart  struct {
		FramePointer bool
		ExternCalls  bool
	}
	StackFrameEnd struct{ FramePointer bool }
	Syscall       struct{}
	Xor           rrr
)