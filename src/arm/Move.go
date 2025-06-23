package arm

import (
	"encoding/binary"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/sizeof"
)

// MoveRegisterNumber moves a number into the given register.
func MoveRegisterNumber(code []byte, destination cpu.Register, number int) []byte {
	instruction, encodable := MoveRegisterNumberSI(destination, number)

	if encodable {
		return binary.LittleEndian.AppendUint32(code, instruction)
	}

	return MoveRegisterNumberMI(code, destination, number)
}

// MoveRegisterNumberMI moves a number into the given register using movz and a series of movk instructions.
func MoveRegisterNumberMI(code []byte, destination cpu.Register, number int) []byte {
	movz := MoveZero(destination, 0, uint16(number))
	code = binary.LittleEndian.AppendUint32(code, movz)
	num := uint64(number)
	halfword := 1

	for {
		num >>= 16

		if num == 0 {
			return code
		}

		movk := MoveKeep(destination, halfword, uint16(num))
		code = binary.LittleEndian.AppendUint32(code, movk)
		halfword++
	}
}

// MoveRegisterNumberSI moves a number into the given register using a single instruction.
func MoveRegisterNumberSI(destination cpu.Register, number int) (uint32, bool) {
	if sizeof.Signed(number) <= 2 {
		if number < 0 {
			return MoveInvertedNumber(destination, uint16(^number), 0), true
		}

		return MoveZero(destination, 0, uint16(number)), true
	}

	if (number&0xFFFFFFFFFFFF == 0xFFFFFFFFFFFF) && sizeof.Signed(number>>48) <= 2 {
		return MoveInvertedNumber(destination, uint16((^number)>>48), 3), true
	}

	code, encodable := MoveBitmaskNumber(destination, number)

	if encodable {
		return code, true
	}

	if (number&0xFFFFFFFF == 0xFFFFFFFF) && sizeof.Signed(number>>32) <= 2 {
		return MoveInvertedNumber(destination, uint16((^number)>>32), 2), true
	}

	if (number&0xFFFF == 0xFFFF) && sizeof.Signed(number>>16) <= 2 {
		return MoveInvertedNumber(destination, uint16((^number)>>16), 1), true
	}

	return 0, false
}

// MoveRegisterRegister copies a register to another register.
func MoveRegisterRegister(destination cpu.Register, source cpu.Register) uint32 {
	if source == SP || destination == SP {
		code, _ := AddRegisterNumber(destination, source, 0)
		return code
	}

	return OrRegisterRegister(destination, ZR, source)
}

// MoveBitmaskNumber moves a bitmask immediate value to a register.
func MoveBitmaskNumber(destination cpu.Register, number int) (uint32, bool) {
	return OrRegisterNumber(destination, ZR, number)
}

// MoveInvertedNumber moves an inverted 16-bit immediate value to a register.
func MoveInvertedNumber(destination cpu.Register, number uint16, shift uint32) uint32 {
	return 0b100100101<<23 | shift<<21 | regImm(destination, number)
}

// MoveKeep moves a 16-bit integer into the given register and keeps all other bits.
func MoveKeep(destination cpu.Register, halfword int, number uint16) uint32 {
	return 0b111100101<<23 | regImmHw(destination, halfword, number)
}

// MoveZero moves a 16-bit integer into the given register and clears all other bits to zero.
func MoveZero(destination cpu.Register, halfword int, number uint16) uint32 {
	return 0b110100101<<23 | regImmHw(destination, halfword, number)
}