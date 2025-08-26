package asm

import (
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/token"
)

type (
	Add        rrr
	AddNumber  rrn
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
	CompareNumber   rn
	Divide          rrr
	Jump            struct {
		Label     string
		Condition token.Kind
	}
	Label struct {
		Name  string
		Align uint8
	}
	Load struct {
		Base        cpu.Register
		Index       cpu.Register
		Destination cpu.Register
		Length      byte
	}
	Modulo    rrr
	Move      rr
	MoveLabel struct {
		Label       string
		Destination cpu.Register
	}
	MoveNumber             rn
	Multiply               rrr
	Negate                 rr
	Or                     rrr
	OrNumber               rrn
	Pop                    struct{ Registers []cpu.Register }
	Push                   struct{ Registers []cpu.Register }
	Return                 struct{}
	ShiftLeft              rrr
	ShiftLeftNumber        rrn
	ShiftRightSigned       rrr
	ShiftRightSignedNumber rrn
	Store                  struct {
		Base   cpu.Register
		Index  cpu.Register
		Source cpu.Register
		Length byte
	}
	StoreNumber struct {
		Base   cpu.Register
		Index  cpu.Register
		Number int
		Length byte
	}
	Subtract        rrr
	SubtractNumber  rrn
	StackFrameStart struct {
		FramePointer bool
		ExternCalls  bool
	}
	StackFrameEnd struct{ FramePointer bool }
	Syscall       struct{}
	Xor           rrr
	XorNumber     rrn
)