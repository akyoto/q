package x86

import "git.urbach.dev/cli/q/src/cpu"

// Call places the return address on the top of the stack and continues
// program flow at the new address.
// The address is relative to the next instruction.
func Call(code []byte, offset uint32) []byte {
	return append(
		code,
		0xE8,
		byte(offset),
		byte(offset>>8),
		byte(offset>>16),
		byte(offset>>24),
	)
}

// Calls a function whose address is stored in the given register.
func CallRegister(code []byte, register cpu.Register) []byte {
	if register > 0b111 {
		code = append(code, 0x41)
		register &= 0b111
	}

	return append(
		code,
		0xFF,
		0xD0+byte(register),
	)
}

// CallAt calls a function at the address stored at the given memory address.
// The memory address is relative to the next instruction.
func CallAt(code []byte, address uint32) []byte {
	return append(
		code,
		0xFF,
		0x15,
		byte(address),
		byte(address>>8),
		byte(address>>16),
		byte(address>>24),
	)
}

// CallAtMemory calls a function at the address stored at the given memory address.
// The memory address is relative to the next instruction.
func CallAtMemory(code []byte, base cpu.Register, offset int8, scale Scale) []byte {
	return memAccess(code, 0xFF, 0xFF, 0b010, base, offset, scale, 4)
}