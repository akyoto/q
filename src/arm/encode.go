package arm

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// memory encodes an instruction with a register, a base register, an addressing mode and an offset.
func memory(destination cpu.Register, base cpu.Register, mode AddressMode, imm9 int) uint32 {
	return uint32(imm9&mask9)<<12 | uint32(mode)<<10 | uint32(base)<<5 | uint32(destination)
}

// pair encodes an instruction using a register pair with memory.
func pair(reg1 cpu.Register, reg2 cpu.Register, base cpu.Register, imm7 int) uint32 {
	return uint32(imm7&mask7)<<15 | uint32(reg2)<<10 | uint32(base)<<5 | uint32(reg1)
}

// regImm encodes an instruction with a register and an immediate.
func regImm(d cpu.Register, imm16 uint16) uint32 {
	return uint32(imm16)<<5 | uint32(d)
}

// regImmHw encodes an instruction with a register, an immediate and
// the 2-bit halfword specifying which 16-bit region of the register is addressed.
func regImmHw(d cpu.Register, hw int, imm16 uint16) uint32 {
	return uint32(hw)<<21 | uint32(imm16)<<5 | uint32(d)
}

// reg2Imm encodes an instruction with 2 registers and an immediate.
func reg2Imm(d cpu.Register, n cpu.Register, imm12 int) uint32 {
	return uint32(imm12&mask12)<<10 | uint32(n)<<5 | uint32(d)
}

// reg2BitmaskImm encodes an instruction with 2 registers and a bitmask immediate.
func reg2BitmaskImm(d cpu.Register, n cpu.Register, N int, immr int, imms int) uint32 {
	return uint32(N)<<22 | uint32(immr)<<16 | uint32(imms)<<10 | uint32(n)<<5 | uint32(d)
}

// reg3 encodes an instruction with 3 registers.
func reg3(d cpu.Register, n cpu.Register, m cpu.Register) uint32 {
	return uint32(m)<<16 | uint32(n)<<5 | uint32(d)
}

// reg3Imm encodes an instruction with 3 registers.
func reg3Imm(d cpu.Register, n cpu.Register, m cpu.Register, imm6 int) uint32 {
	return uint32(m)<<16 | uint32(imm6&mask6)<<10 | uint32(n)<<5 | uint32(d)
}

// reg4 encodes an instruction with 4 registers.
func reg4(d cpu.Register, n cpu.Register, m cpu.Register, a cpu.Register) uint32 {
	return uint32(m)<<16 | uint32(a)<<10 | uint32(n)<<5 | uint32(d)
}

// size encodes the size of the operation.
func size(length byte) uint32 {
	switch length {
	case 1:
		return 0b00
	case 2:
		return 0b01
	case 4:
		return 0b10
	default:
		return 0b11
	}
}