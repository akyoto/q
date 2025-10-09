package asm

import (
	"encoding/binary"
	"fmt"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/token"
)

type compilerARM struct {
	*compiler
}

func (c *compiler) append(code uint32) {
	c.code = binary.LittleEndian.AppendUint32(c.code, code)
}

func (c *compilerARM) Compile(instr Instruction) {
	switch instr := instr.(type) {
	case *Add:
		c.append(arm.AddRegisterRegister(instr.Destination, instr.Source, instr.Operand))
	case *AddNumber:
		code, _ := arm.AddRegisterNumber(instr.Destination, instr.Source, instr.Number)
		c.append(code)
	case *And:
		c.append(arm.AndRegisterRegister(instr.Destination, instr.Source, instr.Operand))
	case *AndNumber:
		code, _ := arm.AndRegisterNumber(instr.Destination, instr.Source, instr.Number)
		c.append(code)
	case *Call:
		c.append(arm.Call(0))
		patch := c.PatchLast4Bytes()

		patch.apply = func(code []byte) []byte {
			address, exists := c.labels[instr.Label]

			if !exists {
				panic("unknown label: " + instr.Label)
			}

			offset := (address - patch.start) / 4
			binary.LittleEndian.PutUint32(code, arm.Call(offset))
			return code
		}
	case *CallExtern:
		encoding, _ := arm.LoadAddress(arm.X16, 0)
		c.append(encoding)
		patch := c.PatchLast4Bytes()
		c.append(arm.LoadFixedOffset(arm.X16, arm.X16, arm.UnscaledImmediate, 0, 8))
		c.append(arm.CallRegister(arm.X16))

		patch.apply = func(code []byte) []byte {
			index := c.libraries.Index(instr.Library, instr.Function)

			if index == -1 {
				panic(fmt.Sprintf("unknown extern function '%s' in library '%s'", instr.Function, instr.Library))
			}

			address := c.importsStart + index*8
			offset := address - patch.start
			encoding, encodable := arm.LoadAddress(arm.X16, offset)

			if !encodable {
				panic("label offset outside of encodable range")
			}

			binary.LittleEndian.PutUint32(code, encoding)
			return code
		}
	case *CallExternStart:
	case *CallExternEnd:
	case *CallRegister:
		c.append(arm.CallRegister(instr.Address))
	case *Compare:
		c.append(arm.CompareRegisterRegister(instr.Destination, instr.Source))
	case *CompareNumber:
		code, _ := arm.CompareRegisterNumber(instr.Destination, instr.Number)
		c.append(code)
	case *Divide:
		c.append(arm.DivUnsignedRegisterRegister(instr.Destination, instr.Source, instr.Operand))
	case *DivideSigned:
		c.append(arm.DivSignedRegisterRegister(instr.Destination, instr.Source, instr.Operand))
	case *Jump:
		c.append(arm.Jump(0))
		patch := c.PatchLast4Bytes()

		patch.apply = func(code []byte) []byte {
			address, exists := c.labels[instr.Label]

			if !exists {
				panic("unknown label: " + instr.Label)
			}

			offset := (address - patch.start) / 4

			var (
				minOffset int
				maxOffset int
			)

			if instr.Condition == token.Invalid {
				minOffset = -(1 << 25)
				maxOffset = (1 << 25) - 1
			} else {
				minOffset = -(1 << 18)
				maxOffset = (1 << 18) - 1
			}

			if offset < minOffset || offset > maxOffset {
				panic("not implemented: long jumps")
			}

			switch instr.Condition {
			case token.Equal:
				binary.LittleEndian.PutUint32(code, arm.JumpIfEqual(offset))
			case token.NotEqual:
				binary.LittleEndian.PutUint32(code, arm.JumpIfNotEqual(offset))
			case token.Greater:
				binary.LittleEndian.PutUint32(code, arm.JumpIfGreater(offset))
			case token.GreaterEqual:
				binary.LittleEndian.PutUint32(code, arm.JumpIfGreaterOrEqual(offset))
			case token.Less:
				binary.LittleEndian.PutUint32(code, arm.JumpIfLess(offset))
			case token.LessEqual:
				binary.LittleEndian.PutUint32(code, arm.JumpIfLessOrEqual(offset))
			default:
				binary.LittleEndian.PutUint32(code, arm.Jump(offset))
			}

			return code
		}
	case *Label:
		c.labels[instr.Name] = len(c.code)
	case *Load:
		scale := arm.Scale1

		if instr.Scale {
			scale = arm.ScaleLength
		}

		c.append(arm.Load(instr.Destination, instr.Base, instr.Index, scale, instr.Length))
	case *LoadFixedOffset:
		if instr.Scale {
			c.append(arm.LoadFixedOffsetScaled(instr.Destination, instr.Base, arm.UnscaledImmediate, uint(instr.Index), instr.Length))
		} else {
			c.append(arm.LoadFixedOffset(instr.Destination, instr.Base, arm.UnscaledImmediate, instr.Index, instr.Length))
		}
	case *Modulo:
		if instr.Destination == instr.Source || instr.Destination == instr.Operand {
			panic("modulo destination register cannot be equal to the source or operand register")
		}

		c.append(arm.DivUnsignedRegisterRegister(instr.Destination, instr.Source, instr.Operand))
		c.append(arm.MultiplySubtract(instr.Destination, instr.Destination, instr.Operand, instr.Source))
	case *ModuloSigned:
		if instr.Destination == instr.Source || instr.Destination == instr.Operand {
			panic("modulo destination register cannot be equal to the source or operand register")
		}

		c.append(arm.DivSignedRegisterRegister(instr.Destination, instr.Source, instr.Operand))
		c.append(arm.MultiplySubtract(instr.Destination, instr.Destination, instr.Operand, instr.Source))
	case *MoveLabel:
		encoding, _ := arm.LoadAddress(instr.Destination, 0)
		c.append(encoding)
		patch := c.PatchLast4Bytes()

		patch.apply = func(code []byte) []byte {
			address, exists := c.labels[instr.Label]

			if !exists {
				panic("unknown label: " + instr.Label)
			}

			offset := address - patch.start
			encoding, encodable := arm.LoadAddress(instr.Destination, offset)

			if !encodable {
				panic("label offset outside of encodable range")
			}

			binary.LittleEndian.PutUint32(code, encoding)
			return code
		}
	case *Move:
		c.append(arm.MoveRegisterRegister(instr.Destination, instr.Source))
	case *MoveNumber:
		c.code = arm.MoveRegisterNumber(c.code, instr.Destination, instr.Number)
	case *Multiply:
		c.append(arm.MulRegisterRegister(instr.Destination, instr.Source, instr.Operand))
	case *Negate:
		c.append(arm.NegateRegister(instr.Destination, instr.Source))
	case *Or:
		c.append(arm.OrRegisterRegister(instr.Destination, instr.Source, instr.Operand))
	case *OrNumber:
		code, _ := arm.OrRegisterNumber(instr.Destination, instr.Source, instr.Number)
		c.append(code)
	case *Pop:
		registers := instr.Registers
		count := len(registers)

		if count&1 != 0 {
			count--
			c.append(arm.LoadFixedOffset(registers[count], arm.SP, arm.PostIndex, 16, 8))
		}

		for i := count - 2; i >= 0; i -= 2 {
			c.append(arm.LoadPair(registers[i], registers[i+1], arm.SP, 16))
		}
	case *Push:
		registers := instr.Registers

		for i := 0; i < len(registers); i += 2 {
			if i+1 < len(registers) {
				c.append(arm.StorePair(registers[i], registers[i+1], arm.SP, -16))
			} else {
				c.append(arm.StoreFixedOffsetRegister(registers[i], arm.SP, arm.PreIndex, -16, 8))
			}
		}
	case *Return:
		c.append(arm.Return())
	case *ShiftLeft:
		c.append(arm.ShiftLeft(instr.Destination, instr.Source, instr.Operand))
	case *ShiftLeftNumber:
		c.append(arm.ShiftLeftNumber(instr.Destination, instr.Source, instr.Number))
	case *ShiftRight:
		c.append(arm.ShiftRight(instr.Destination, instr.Source, instr.Operand))
	case *ShiftRightNumber:
		c.append(arm.ShiftRightNumber(instr.Destination, instr.Source, instr.Number))
	case *ShiftRightSigned:
		c.append(arm.ShiftRightSigned(instr.Destination, instr.Source, instr.Operand))
	case *ShiftRightSignedNumber:
		c.append(arm.ShiftRightSignedNumber(instr.Destination, instr.Source, instr.Number))
	case *Store:
		scale := arm.Scale1

		if instr.Scale {
			scale = arm.ScaleLength
		}

		c.append(arm.StoreRegister(instr.Source, instr.Base, instr.Index, scale, instr.Length))
	case *StoreNumber:
		panic("arm64 does not support memory stores of immediates")
	case *Subtract:
		c.append(arm.SubRegisterRegister(instr.Destination, instr.Source, instr.Operand))
	case *SubtractNumber:
		code, _ := arm.SubRegisterNumber(instr.Destination, instr.Source, instr.Number)
		c.append(code)
	case *StackFrameStart:
		c.append(arm.StoreFixedOffsetRegister(arm.LR, arm.SP, arm.PreIndex, -16, 8))
	case *StackFrameEnd:
		c.append(arm.LoadFixedOffset(arm.LR, arm.SP, arm.PostIndex, 16, 8))
	case *Syscall:
		switch c.build.OS {
		case config.Mac:
			c.append(arm.Syscall(0x80))
		default:
			c.append(arm.Syscall(0))
		}
	case *Xor:
		c.append(arm.XorRegisterRegister(instr.Destination, instr.Source, instr.Operand))
	case *XorNumber:
		code, _ := arm.XorRegisterNumber(instr.Destination, instr.Source, instr.Number)
		c.append(code)
	default:
		panic("unknown instruction")
	}
}