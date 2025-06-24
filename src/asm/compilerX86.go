package asm

import (
	"encoding/binary"

	"git.urbach.dev/cli/q/src/x86"
)

type compilerX86 struct {
	*compiler
}

func (c *compilerX86) Compile(instr Instruction) {
	switch instr := instr.(type) {
	case *Call:
		c.code = x86.Call(c.code, 0)
		end := len(c.code)

		c.Defer(func() {
			address, exists := c.labels[instr.Label]

			if !exists {
				panic("unknown label: " + instr.Label)
			}

			offset := address - end
			binary.LittleEndian.PutUint32(c.code[end-4:end], uint32(offset))
		})
	case *Label:
		c.labels[instr.Name] = len(c.code)
	case *MoveRegisterLabel:
		c.code = x86.LoadAddress(c.code, instr.Destination, 0)
		end := len(c.code)

		c.Defer(func() {
			address, exists := c.labels[instr.Label]

			if !exists {
				panic("unknown label: " + instr.Label)
			}

			offset := address - end
			binary.LittleEndian.PutUint32(c.code[end-4:end], uint32(offset))
		})
	case *MoveRegisterNumber:
		c.code = x86.MoveRegisterNumber(c.code, instr.Destination, instr.Number)
	case *MoveRegisterRegister:
		c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
	case *Return:
		c.code = x86.Return(c.code)
	case *Syscall:
		c.code = x86.Syscall(c.code)
	}
}