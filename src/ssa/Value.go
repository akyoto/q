package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/cpu"
)

// Value is a single instruction in a basic block.
// It is implemented as a "fat struct" for performance reasons.
// It contains all the fields necessary to represent all instruction types.
type Value struct {
	Args     []*Value
	Int      int
	Text     string
	Register cpu.Register
	Type     Type
}

// Equals returns true if the values are equal.
func (a Value) Equals(b Value) bool {
	if a.Type != b.Type {
		return false
	}

	if a.Int != b.Int {
		return false
	}

	if a.Text != b.Text {
		return false
	}

	if a.Register != b.Register {
		return false
	}

	if len(a.Args) != len(b.Args) {
		return false
	}

	for i := range a.Args {
		if !a.Args[i].Equals(*b.Args[i]) {
			return false
		}
	}

	return true
}

// IsConst returns true if the value is constant.
func (i *Value) IsConst() bool {
	switch i.Type {
	case Func, Int, Register, String:
		return true
	default:
		return false
	}
}

// String returns a human-readable representation of the instruction.
func (i *Value) String() string {
	switch i.Type {
	case Func:
		return i.Text
	case Int:
		return fmt.Sprintf("%d", i.Int)
	case Register:
		return i.Register.String()
	case String:
		return fmt.Sprintf("\"%s\"", i.Text)
	case Add:
		return fmt.Sprintf("%s + %s", i.Args[0], i.Args[1])
	case Return:
		return fmt.Sprintf("return %s", i.Args[0])
	case Call:
		return fmt.Sprintf("call%v", i.Args)
	case Syscall:
		return fmt.Sprintf("syscall%v", i.Args)
	default:
		return ""
	}
}