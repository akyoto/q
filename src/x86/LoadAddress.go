package x86

import (
	"encoding/binary"

	"git.urbach.dev/cli/q/src/cpu"
)

// LoadAddress calculates the address with the RIP-relative offset and writes the result to the destination register.
func LoadAddress(code []byte, destination cpu.Register, offset int) []byte {
	code = encode(code, AddressMemory, destination, 0b101, 8, 0x8D)
	return binary.LittleEndian.AppendUint32(code, uint32(offset))
}