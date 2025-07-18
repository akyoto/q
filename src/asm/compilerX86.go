package asm

import (
	"encoding/binary"
	"fmt"
	"slices"

	"git.urbach.dev/cli/q/src/sizeof"
	"git.urbach.dev/cli/q/src/token"
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
		patch := c.PatchLast4Bytes()

		patch.apply = func(code []byte) []byte {
			address, exists := c.labels[instr.Label]

			if !exists {
				panic("unknown label: " + instr.Label)
			}

			offset := address - patch.end
			binary.LittleEndian.PutUint32(code, uint32(offset))
			return code
		}
	case *CallExtern:
		c.code = x86.CallAt(c.code, 0)
		patch := c.PatchLast4Bytes()

		patch.apply = func(code []byte) []byte {
			index := c.libraries.Index(instr.Library, instr.Function)

			if index == -1 {
				panic(fmt.Sprintf("unknown extern function '%s' in library '%s'", instr.Function, instr.Library))
			}

			address := c.importsStart + index*8
			offset := address - patch.end
			binary.LittleEndian.PutUint32(code, uint32(offset))
			return code
		}
	case *CompareRegisterRegister:
		c.code = x86.CompareRegisterRegister(c.code, instr.SourceA, instr.SourceB)
	case *DivRegisterRegister:
		if instr.Source != x86.R0 {
			c.code = x86.MoveRegisterRegister(c.code, x86.R0, instr.Source)
		}

		c.code = x86.ExtendR0ToR2(c.code)
		c.code = x86.DivRegister(c.code, instr.Operand)

		if instr.Destination != x86.R0 {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, x86.R0)
		}
	case *Jump:
		switch instr.Condition {
		case token.Equal:
			c.code = x86.Jump8IfEqual(c.code, 0x00)
		case token.NotEqual:
			c.code = x86.Jump8IfNotEqual(c.code, 0x00)
		case token.Greater:
			c.code = x86.Jump8IfGreater(c.code, 0x00)
		case token.GreaterEqual:
			c.code = x86.Jump8IfGreaterOrEqual(c.code, 0x00)
		case token.Less:
			c.code = x86.Jump8IfLess(c.code, 0x00)
		case token.LessEqual:
			c.code = x86.Jump8IfLessOrEqual(c.code, 0x00)
		default:
			c.code = x86.Jump8(c.code, 0x00)
		}

		patch := &patch{
			start: len(c.code) - 2,
			end:   len(c.code),
		}

		patch.apply = func(code []byte) []byte {
			address, exists := c.labels[instr.Label]

			if !exists {
				panic("unknown label: " + instr.Label)
			}

			offset := address - patch.end

			switch code[0] {
			case 0x74, 0x75, 0x7C, 0x7D, 0x7E, 0x7F, 0xEB:
				if sizeof.Signed(offset) == 1 {
					code[1] = byte(offset)
					return code
				}

				var jump []byte

				switch code[0] {
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
					panic(fmt.Sprintf("failed to increase pointer size for instruction 0x%x", code[0]))
				}

				shift := len(jump) + 2
				offset -= shift
				jump = binary.LittleEndian.AppendUint32(jump, uint32(offset))
				return jump
			case 0xE9:
				binary.LittleEndian.PutUint32(code[1:], uint32(offset))
				return code
			case 0x0F:
				binary.LittleEndian.PutUint32(code[2:], uint32(offset))
				return code
			default:
				panic(fmt.Sprintf("not a jump instruction 0x%x", code[0]))
			}
		}

		c.earlyPatches = append(c.earlyPatches, patch)
	case *Label:
		c.labels[instr.Name] = len(c.code)
	case *MoveRegisterLabel:
		c.code = x86.LoadAddress(c.code, instr.Destination, 0)
		patch := c.PatchLast4Bytes()

		patch.apply = func(code []byte) []byte {
			address, exists := c.labels[instr.Label]

			if !exists {
				panic("unknown label: " + instr.Label)
			}

			offset := address - patch.end
			binary.LittleEndian.PutUint32(code, uint32(offset))
			return code
		}
	case *MoveRegisterNumber:
		c.code = x86.MoveRegisterNumber(c.code, instr.Destination, instr.Number)
	case *MoveRegisterRegister:
		c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
	case *MulRegisterRegister:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.MulRegisterRegister(c.code, instr.Destination, instr.Operand)
	case *PopRegisters:
		for _, register := range slices.Backward(instr.Registers) {
			c.code = x86.PopRegister(c.code, register)
		}
	case *PushRegisters:
		for _, register := range instr.Registers {
			c.code = x86.PushRegister(c.code, register)
		}
	case *Return:
		c.code = x86.Return(c.code)
	case *SubRegisterNumber:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.SubRegisterNumber(c.code, instr.Destination, instr.Number)
	case *SubRegisterRegister:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.SubRegisterRegister(c.code, instr.Destination, instr.Operand)
	case *StackFrameStart:
		if instr.FramePointer {
			c.code = x86.PushRegister(c.code, x86.R5)
			c.code = x86.MoveRegisterRegister(c.code, x86.R5, x86.SP)
		}

		if instr.ExternCalls {
			c.code = x86.AndRegisterNumber(c.code, x86.SP, -16)
			c.code = x86.SubRegisterNumber(c.code, x86.SP, 32)
		}
	case *StackFrameEnd:
		if instr.FramePointer {
			c.code = x86.MoveRegisterRegister(c.code, x86.SP, x86.R5)
			c.code = x86.PopRegister(c.code, x86.R5)
		}
	case *Syscall:
		c.code = x86.Syscall(c.code)
	default:
		panic("unknown instruction")
	}
}