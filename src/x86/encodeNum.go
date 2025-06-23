package x86

import (
	"encoding/binary"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/sizeof"
)

// encodeNum encodes an instruction with up to two registers and a number parameter.
func encodeNum(code []byte, mod AddressMode, reg cpu.Register, rm cpu.Register, number int, opCode8 byte, opCode32 byte) []byte {
	if sizeof.Signed(int64(number)) == 1 {
		code = encode(code, mod, reg, rm, 8, opCode8)
		return append(code, byte(number))
	}

	code = encode(code, mod, reg, rm, 8, opCode32)
	return binary.LittleEndian.AppendUint32(code, uint32(number))
}