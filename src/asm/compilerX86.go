package asm

import (
	"encoding/binary"
	"fmt"
	"slices"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/x86"
)

type compilerX86 struct {
	*compiler
}

func (c *compilerX86) Compile(instr Instruction) {
	switch instr := instr.(type) {
	case *Add:
		if instr.Destination == instr.Operand {
			panic("add destination register cannot be equal to the operand register")
		}

		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.AddRegisterRegister(c.code, instr.Destination, instr.Operand)
	case *AddNumber:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.AddRegisterNumber(c.code, instr.Destination, instr.Number)
	case *And:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.AndRegisterRegister(c.code, instr.Destination, instr.Operand)
	case *AndNumber:
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
		c.code = x86.SubRegisterNumber(c.code, x86.SP, 32)
		c.code = x86.CallAt(c.code, 0)
		patch := c.PatchLast4Bytes()
		c.code = x86.AddRegisterNumber(c.code, x86.SP, 32)

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
	case *CallRegister:
		c.code = x86.CallRegister(c.code, instr.Address)
	case *Compare:
		c.code = x86.CompareRegisterRegister(c.code, instr.Destination, instr.Source)
	case *CompareNumber:
		if instr.Number == 0 {
			c.code = x86.TestRegisterRegister(c.code, instr.Destination, instr.Destination)
		} else {
			c.code = x86.CompareRegisterNumber(c.code, instr.Destination, instr.Number)
		}
	case *Divide:
		if instr.Operand == x86.R2 {
			panic("divisor register cannot be R2")
		}

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
				if cpu.SizeInt(offset) == 1 {
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
		if instr.Align > 0 {
			_, pad := exe.AlignPad(len(c.code), int(instr.Align))
			c.code = x86.Nop(c.code, pad)
		}

		c.labels[instr.Name] = len(c.code)
	case *Load:
		scale := toX86Scale(instr.Scale, instr.Length)

		if instr.Length <= 2 {
			c.code = x86.LoadDynamicRegisterZeroExtend(c.code, instr.Destination, instr.Base, instr.Index, scale, instr.Length)
		} else {
			c.code = x86.LoadDynamicRegister(c.code, instr.Destination, instr.Base, instr.Index, scale, instr.Length)
		}
	case *Modulo:
		if instr.Operand == x86.R0 {
			panic("modulo operand register cannot be R0")
		}

		if instr.Source != x86.R0 {
			c.code = x86.MoveRegisterRegister(c.code, x86.R0, instr.Source)
		}

		c.code = x86.ExtendR0ToR2(c.code)
		c.code = x86.DivRegister(c.code, instr.Operand)

		if instr.Destination != x86.R2 {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, x86.R2)
		}
	case *MoveLabel:
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
	case *Move:
		c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
	case *MoveNumber:
		c.code = x86.MoveRegisterNumber(c.code, instr.Destination, instr.Number)
	case *Multiply:
		if instr.Destination == instr.Operand {
			panic("multiply destination register cannot be equal to the operand register")
		}

		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.MulRegisterRegister(c.code, instr.Destination, instr.Operand)
	case *Negate:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.NegateRegister(c.code, instr.Destination)
	case *Or:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.OrRegisterRegister(c.code, instr.Destination, instr.Operand)
	case *OrNumber:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.OrRegisterNumber(c.code, instr.Destination, instr.Number)
	case *Pop:
		for _, register := range slices.Backward(instr.Registers) {
			c.code = x86.PopRegister(c.code, register)
		}
	case *Push:
		for _, register := range instr.Registers {
			c.code = x86.PushRegister(c.code, register)
		}
	case *Return:
		c.code = x86.Return(c.code)
	case *ShiftLeft:
		c.code = prepareShiftX86(c.code, instr.Destination, instr.Source, instr.Operand)
		c.code = x86.ShiftLeft(c.code, instr.Destination)
	case *ShiftLeftNumber:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.ShiftLeftNumber(c.code, instr.Destination, byte(instr.Number))
	case *ShiftRight:
		c.code = prepareShiftX86(c.code, instr.Destination, instr.Source, instr.Operand)
		c.code = x86.ShiftRight(c.code, instr.Destination)
	case *ShiftRightNumber:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.ShiftRightNumber(c.code, instr.Destination, byte(instr.Number))
	case *ShiftRightSigned:
		c.code = prepareShiftX86(c.code, instr.Destination, instr.Source, instr.Operand)
		c.code = x86.ShiftRightSigned(c.code, instr.Destination)
	case *ShiftRightSignedNumber:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.ShiftRightSignedNumber(c.code, instr.Destination, byte(instr.Number))
	case *StackFrameStart:
		if instr.FramePointer {
			c.code = x86.PushRegister(c.code, x86.R5)
			c.code = x86.MoveRegisterRegister(c.code, x86.R5, x86.SP)
		}

		if instr.ExternCalls {
			c.code = x86.AndRegisterNumber(c.code, x86.SP, -16)
		}
	case *StackFrameEnd:
		if instr.FramePointer {
			c.code = x86.MoveRegisterRegister(c.code, x86.SP, x86.R5)
			c.code = x86.PopRegister(c.code, x86.R5)
		}
	case *Store:
		scale := toX86Scale(instr.Scale, instr.Length)
		c.code = x86.StoreDynamicRegister(c.code, instr.Base, instr.Index, scale, instr.Length, instr.Source)
	case *StoreNumber:
		scale := toX86Scale(instr.Scale, instr.Length)
		c.code = x86.StoreDynamicNumber(c.code, instr.Base, instr.Index, scale, instr.Length, instr.Number)
	case *Subtract:
		if instr.Destination == instr.Operand {
			panic("subtract destination register cannot be equal to the operand register")
		}

		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.SubRegisterRegister(c.code, instr.Destination, instr.Operand)
	case *SubtractNumber:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.SubRegisterNumber(c.code, instr.Destination, instr.Number)
	case *Syscall:
		c.code = x86.Syscall(c.code)
	case *Xor:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.XorRegisterRegister(c.code, instr.Destination, instr.Operand)
	case *XorNumber:
		if instr.Destination != instr.Source {
			c.code = x86.MoveRegisterRegister(c.code, instr.Destination, instr.Source)
		}

		c.code = x86.XorRegisterNumber(c.code, instr.Destination, instr.Number)
	default:
		panic("unknown instruction")
	}
}

// prepareShiftX86 checks that the registers are correct for the shift instruction
// and also moves source to destination and the operand to R1 if needed.
func prepareShiftX86(code []byte, destination cpu.Register, source cpu.Register, operand cpu.Register) []byte {
	if destination == x86.R1 {
		panic("shift destination cannot be R1")
	}

	if destination == operand {
		panic("shift destination register cannot be equal to the operand register")
	}

	if destination != source {
		code = x86.MoveRegisterRegister(code, destination, source)
	}

	if operand != x86.R1 {
		code = x86.MoveRegisterRegister(code, x86.R1, operand)
	}

	return code
}

// toX86Scale returns the scale factor for the memory instruction.
func toX86Scale(enable bool, length byte) x86.Scale {
	if !enable {
		return x86.Scale1
	}

	switch length {
	case 1:
		return x86.Scale1
	case 2:
		return x86.Scale2
	case 4:
		return x86.Scale4
	case 8:
		return x86.Scale8
	default:
		panic("unsupported scale")
	}
}