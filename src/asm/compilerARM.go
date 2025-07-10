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
		c.append(arm.LoadRegister(arm.X0, arm.X0, 0, 8))
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
	case *FunctionStart:
		c.append(arm.StorePair(arm.FP, arm.LR, arm.SP, -16))
		c.append(arm.MoveRegisterRegister(arm.FP, arm.SP))
	case *FunctionEnd:
		c.append(arm.LoadPair(arm.FP, arm.LR, arm.SP, 16))
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
	case *PushRegister:
		panic("not implemented")
	case *Return:
		c.append(arm.Return())
	case *Syscall:
		c.append(arm.Syscall())
	default:
		panic("unknown instruction")
	}
}