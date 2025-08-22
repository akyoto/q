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
		c.append(arm.LoadAddress(arm.X16, 0))
		patch := c.PatchLast4Bytes()
		c.append(arm.LoadRegister(arm.X16, arm.X16, arm.UnscaledImmediate, 0, 8))
		c.append(arm.CallRegister(arm.X16))

		patch.apply = func(code []byte) []byte {
			index := c.libraries.Index(instr.Library, instr.Function)

			if index == -1 {
				panic(fmt.Sprintf("unknown extern function '%s' in library '%s'", instr.Function, instr.Library))
			}

			address := c.importsStart + index*8
			offset := address - patch.start
			binary.LittleEndian.PutUint32(code, arm.LoadAddress(arm.X16, offset))
			return code
		}
	case *CallExternStart:
	case *CallExternEnd:
	case *Compare:
		c.append(arm.CompareRegisterRegister(instr.Destination, instr.Source))
	case *CompareNumber:
		code, _ := arm.CompareRegisterNumber(instr.Destination, instr.Number)
		c.append(code)
	case *Divide:
		c.append(arm.DivRegisterRegister(instr.Destination, instr.Source, instr.Operand))
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
		c.append(arm.LoadDynamicRegister(instr.Destination, instr.Base, arm.Offset, instr.Index, instr.Length))
	case *Modulo:
		if instr.Destination == instr.Source || instr.Destination == instr.Operand {
			panic("modulo operation needs a separate destination register")
		}

		c.append(arm.DivRegisterRegister(instr.Destination, instr.Source, instr.Operand))
		c.append(arm.MultiplySubtract(instr.Destination, instr.Destination, instr.Operand, instr.Source))
	case *MoveLabel:
		c.append(arm.LoadAddress(instr.Destination, 0))
		patch := c.PatchLast4Bytes()

		patch.apply = func(code []byte) []byte {
			address, exists := c.labels[instr.Label]

			if !exists {
				panic("unknown label: " + instr.Label)
			}

			offset := address - patch.start
			binary.LittleEndian.PutUint32(code, arm.LoadAddress(instr.Destination, offset))
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
	case *Pop:
		registers := instr.Registers
		count := len(registers)

		if count&1 != 0 {
			count--
			c.append(arm.LoadRegister(registers[count], arm.SP, arm.PostIndex, 16, 8))
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
				c.append(arm.StoreRegister(registers[i], arm.SP, arm.PreIndex, -16, 8))
			}
		}
	case *Return:
		c.append(arm.Return())
	case *ShiftLeft:
		c.append(arm.ShiftLeft(instr.Destination, instr.Source, instr.Operand))
	case *ShiftLeftNumber:
		c.append(arm.ShiftLeftNumber(instr.Destination, instr.Source, instr.Number))
	case *ShiftRightSigned:
		c.append(arm.ShiftRightSigned(instr.Destination, instr.Source, instr.Operand))
	case *ShiftRightSignedNumber:
		c.append(arm.ShiftRightSignedNumber(instr.Destination, instr.Source, instr.Number))
	case *Store:
		c.append(arm.StoreDynamicRegister(instr.Source, instr.Base, arm.Offset, instr.Index, instr.Length))
	case *Subtract:
		c.append(arm.SubRegisterRegister(instr.Destination, instr.Source, instr.Operand))
	case *SubtractNumber:
		code, _ := arm.SubRegisterNumber(instr.Destination, instr.Source, instr.Number)
		c.append(code)
	case *StackFrameStart:
		c.append(arm.StoreRegister(arm.LR, arm.SP, arm.PreIndex, -16, 8))
	case *StackFrameEnd:
		c.append(arm.LoadRegister(arm.LR, arm.SP, arm.PostIndex, 16, 8))
	case *Syscall:
		switch c.build.OS {
		case config.Mac:
			c.append(arm.Syscall(0x80))
		default:
			c.append(arm.Syscall(0))
		}
	case *Xor:
		c.append(arm.XorRegisterRegister(instr.Destination, instr.Source, instr.Operand))
	default:
		panic("unknown instruction")
	}
}