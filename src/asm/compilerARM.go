package asm

import (
	"encoding/binary"
	"fmt"

	"git.urbach.dev/cli/q/src/arm"
)

type compilerARM struct {
	*compiler
}

func (c *compiler) append(code uint32) {
	c.code = binary.LittleEndian.AppendUint32(c.code, code)
}

func (c *compilerARM) Compile(instr Instruction) {
	switch instr := instr.(type) {
	case *AddRegisterRegister:
		c.append(arm.AddRegisterRegister(instr.Destination, instr.Source, instr.Operand))
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
		c.append(arm.LoadAddress(arm.X0, 0))
		patch := c.PatchLast4Bytes()
		c.append(arm.LoadRegister(arm.X0, arm.X0, arm.UnscaledImmediate, 0, 8))
		c.append(arm.CallRegister(arm.X0))

		patch.apply = func(code []byte) []byte {
			index := c.libraries.Index(instr.Library, instr.Function)

			if index == -1 {
				panic(fmt.Sprintf("unknown extern function '%s' in library '%s'", instr.Function, instr.Library))
			}

			address := c.importsStart + index*8
			offset := address - patch.start
			binary.LittleEndian.PutUint32(code, arm.LoadAddress(arm.X0, offset))
			return code
		}
	case *CallExternStart:
	case *CallExternEnd:
	case *DivRegisterRegister:
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

			if offset != (offset & 0b11_11111111_11111111_11111111) {
				panic("not implemented: long jumps")
			}

			binary.LittleEndian.PutUint32(code, arm.Jump(offset))
			return code
		}
	case *Label:
		c.labels[instr.Name] = len(c.code)
	case *MoveRegisterLabel:
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
	case *MoveRegisterNumber:
		c.code = arm.MoveRegisterNumber(c.code, instr.Destination, instr.Number)
	case *MoveRegisterRegister:
		c.append(arm.MoveRegisterRegister(instr.Destination, instr.Source))
	case *MulRegisterRegister:
		c.append(arm.MulRegisterRegister(instr.Destination, instr.Source, instr.Operand))
	case *PopRegisters:
		registers := instr.Registers
		count := len(registers)

		if count&1 != 0 {
			count--
			c.append(arm.LoadRegister(registers[count], arm.SP, arm.PostIndex, 16, 8))
		}

		for i := count - 2; i >= 0; i -= 2 {
			c.append(arm.LoadPair(registers[i], registers[i+1], arm.SP, 16))
		}
	case *PushRegisters:
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
	case *SubRegisterRegister:
		c.append(arm.SubRegisterRegister(instr.Destination, instr.Source, instr.Operand))
	case *StackFrameStart:
		c.append(arm.StoreRegister(arm.LR, arm.SP, arm.PreIndex, -16, 8))
	case *StackFrameEnd:
		c.append(arm.LoadRegister(arm.LR, arm.SP, arm.PostIndex, 16, 8))
	case *Syscall:
		c.append(arm.Syscall())
	default:
		panic("unknown instruction")
	}
}