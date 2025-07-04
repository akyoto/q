package asm

import (
	"encoding/binary"
	"fmt"

	"git.urbach.dev/cli/q/src/sizeof"
	"git.urbach.dev/cli/q/src/x86"
)

type compilerX86 struct {
	*compiler
}

func (c *compilerX86) Compile(instr Instruction) {
	switch instr := instr.(type) {
	case *AddRegisterRegister:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.AddRegisterRegister(c.code, instr.Destination, instr.Operand)
	case *AndRegisterNumber:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.AndRegisterNumber(c.code, instr.Destination, instr.Number)
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
	case *CallExtern:
		c.code = x86.MoveRegisterRegister(c.code, x86.R5, x86.SP)
		c.code = x86.AndRegisterNumber(c.code, x86.SP, -16)
		c.code = x86.SubRegisterNumber(c.code, x86.SP, 32)
		c.code = x86.CallAt(c.code, 0)
		end := len(c.code)
		c.code = x86.MoveRegisterRegister(c.code, x86.SP, x86.R5)

		c.Defer(func() {
			index := c.libraries.Index(instr.Library, instr.Function)

			if index == -1 {
				panic(fmt.Sprintf("unknown extern function '%s' in library '%s'", instr.Function, instr.Library))
			}

			address := c.importsStart + index*8
			offset := address - end
			binary.LittleEndian.PutUint32(c.code[end-4:end], uint32(offset))
		})
	case *FunctionStart:
	case *FunctionEnd:
	case *Jump:
		c.code = x86.Jump8(c.code, 0)
		end := len(c.code)

		c.Defer(func() {
			address, exists := c.labels[instr.Label]

			if !exists {
				panic("unknown label: " + instr.Label)
			}

			offset := address - end

			if sizeof.Signed(offset) > 1 {
				panic("not implemented: long jumps")
			}

			c.code[end-1] = byte(offset)
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
	case *SubRegisterNumber:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.SubRegisterNumber(c.code, instr.Destination, instr.Number)
	case *Syscall:
		c.code = x86.Syscall(c.code)
	default:
		panic("unknown instruction")
	}
}