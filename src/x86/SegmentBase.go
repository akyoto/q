package x86

import "git.urbach.dev/cli/q/src/cpu"

// Gets FS segment base address and writes it to the register.
func ReadFS(code []byte, register cpu.Register) []byte {
	return encode(code, AddressDirect, 0b000, register, 8, 0xF3000FAE)
}

// Gets GS segment base address and writes it to the register.
func ReadGS(code []byte, register cpu.Register) []byte {
	return encode(code, AddressDirect, 0b001, register, 8, 0xF3000FAE)
}

// Sets FS segment base address to the register contents.
func WriteFS(code []byte, register cpu.Register) []byte {
	return encode(code, AddressDirect, 0b010, register, 8, 0xF3000FAE)
}

// Sets GS segment base address to the register contents.
func WriteGS(code []byte, register cpu.Register) []byte {
	return encode(code, AddressDirect, 0b011, register, 8, 0xF3000FAE)
}