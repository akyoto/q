package asm

import (
	"encoding/binary"
	"fmt"
	"slices"

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

		c.Defer(end, func(end int) {
			address, exists := c.labels[instr.Label]

			if !exists {
				panic("unknown label: " + instr.Label)
			}

			offset := address - end
			binary.LittleEndian.PutUint32(c.code[end-4:end], uint32(offset))
		})
	case *CallExtern:
		c.code = x86.CallAt(c.code, 0)
		end := len(c.code)

		c.Defer(end, func(end int) {
			index := c.libraries.Index(instr.Library, instr.Function)

			if index == -1 {
				panic(fmt.Sprintf("unknown extern function '%s' in library '%s'", instr.Function, instr.Library))
			}

			address := c.importsStart + index*8
			offset := address - end
			binary.LittleEndian.PutUint32(c.code[end-4:end], uint32(offset))
		})
	case *CallExternStart:
		c.code = x86.MoveRegisterRegister(c.code, x86.R5, x86.SP)
		c.code = x86.AndRegisterNumber(c.code, x86.SP, -16)
		c.code = x86.SubRegisterNumber(c.code, x86.SP, 32)
	case *CallExternEnd:
		c.code = x86.MoveRegisterRegister(c.code, x86.SP, x86.R5)
	case *FunctionStart:
	case *FunctionEnd:
	case *Jump:
		start := len(c.code)
		c.code = x86.Jump8(c.code, 0)

		c.DeferCodeChange(start, func(start int) bool {
			address, exists := c.labels[instr.Label]

			if !exists {
				panic("unknown label: " + instr.Label)
			}

			var (
				end    int
				offset int
			)

			switch c.code[start] {
			case 0x74, 0x75, 0x7C, 0x7D, 0x7E, 0x7F, 0xEB:
				end = start + 2
				offset = address - end

				if sizeof.Signed(offset) == 1 {
					c.code[end-1] = byte(offset)
					return false
				}
			case 0xE9:
				end = start + 5
				offset = address - end
				binary.LittleEndian.PutUint32(c.code[start+1:end], uint32(offset))
				return false
			case 0x0F:
				end = start + 6
				offset = address - end
				binary.LittleEndian.PutUint32(c.code[start+2:end], uint32(offset))
				return false
			default:
				return false
			}

			var jump []byte

			switch c.code[start] {
			case 0x74: // JE
				jump = []byte{0x0F, 0x84}
			case 0x75: // JNE
				jump = []byte{0x0F, 0x85}
			case 0x7C: // JL
				jump = []byte{0x0F, 0x8C}
			case 0x7D: // JGE
				jump = []byte{0x0F, 0x8D}
			case 0x7E: // JLE
				jump = []byte{0x0F, 0x8E}
			case 0x7F: // JG
				jump = []byte{0x0F, 0x8F}
			case 0xEB: // JMP
				jump = []byte{0xE9}
			default:
				return false
			}

			oldSize := 2
			newSize := len(jump) + 4
			shift := newSize - oldSize
			offset -= shift
			jump = binary.LittleEndian.AppendUint32(jump, uint32(offset))

			for key, address := range c.labels {
				if address >= end {
					c.labels[key] += shift
				}
			}

			for address, call := range c.deferred {
				if address >= end {
					delete(c.deferred, address)
					address += shift
					c.deferred[address] = call
				}
			}

			for address, call := range c.deferredCodeChanges {
				if address >= end {
					delete(c.deferredCodeChanges, address)
					address += shift
					c.deferredCodeChanges[address] = call
				}
			}

			left := c.code[:start]
			right := c.code[end:]
			c.code = slices.Concat(left, jump, right)
			return true
		})
	case *Label:
		c.labels[instr.Name] = len(c.code)
	case *MoveRegisterLabel:
		c.code = x86.LoadAddress(c.code, instr.Destination, 0)
		end := len(c.code)

		c.Defer(end, func(end int) {
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
	case *PushRegister:
		c.code = x86.PushRegister(c.code, instr.Register)
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