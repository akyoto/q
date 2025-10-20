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
	CallRegister    struct{ Address cpu.Register }
	Compare         rr
	CompareNumber   rn
	Divide          rrr
	DivideSigned    rrr
	Jump            struct {
		Label     string
		Condition token.Kind
	}
	Label struct {
		Name  string
		Align uint8
	}
	Load struct {
		Index       cpu.Register
		Base        cpu.Register
		Destination cpu.Register
		Length      byte
		Scale       bool
	}
	LoadFixedOffset struct {
		Index       int
		Base        cpu.Register
		Destination cpu.Register
		Length      byte
		Scale       bool
	}
	Modulo       rrr
	ModuloSigned rrr
	Move         rr
	MoveLabel    struct {
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
	ShiftRight             rrr
	ShiftRightNumber       rrn
	ShiftRightSigned       rrr
	ShiftRightSignedNumber rrn
	Store                  struct {
		Index  cpu.Register
		Base   cpu.Register
		Source cpu.Register
		Length byte
		Scale  bool
	}
	StoreFixedOffset struct {
		Index  int
		Base   cpu.Register
		Source cpu.Register
		Length byte
		Scale  bool
	}
	StoreFixedOffsetNumber struct {
		Index  int
		Number int
		Base   cpu.Register
		Length byte
		Scale  bool
	}
	StoreNumber struct {
		Number int
		Base   cpu.Register
		Index  cpu.Register
		Length byte
		Scale  bool
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