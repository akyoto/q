package x86

import (
	"encoding/binary"

	"git.urbach.dev/cli/q/src/cpu"
)

// MoveRegisterNumber moves an integer into the given register.
func MoveRegisterNumber(code []byte, destination cpu.Register, number int) []byte {
	w := byte(0)

	if cpu.SizeInt(int64(number)) == 8 {
		w = 1
	}

	if w == 0 && number < 0 {
		return moveRegisterNumber32(code, destination, number)
	}

	b := byte(0)

	if destination > 0b111 {
		b = 1
		destination &= 0b111
	}

	if w != 0 || b != 0 {
		code = append(code, REX(w, 0, 0, b))
	}

	code = append(code, 0xB8+byte(destination))

	if w == 1 {
		return binary.LittleEndian.AppendUint64(code, uint64(number))
	}

	return binary.LittleEndian.AppendUint32(code, uint32(number))
}

// moveRegisterNumber32 moves an integer into the given register and sign-extends the register.
func moveRegisterNumber32(code []byte, destination cpu.Register, number int) []byte {
	code = encode(code, AddressDirect, 0, destination, 8, 0xC7)
	return binary.LittleEndian.AppendUint32(code, uint32(number))
}

// MoveRegisterRegister copies a register to another register.
func MoveRegisterRegister(code []byte, destination cpu.Register, source cpu.Register) []byte {
	return encode(code, AddressDirect, source, destination, 8, 0x89)
}