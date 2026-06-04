package x86

import "encoding/binary"

// appendNumber appends the number for memory store instructions.
func appendNumber(code []byte, length byte, number int) []byte {
	switch length {
	case 1:
		return append(code, uint8(number))
	case 2:
		return binary.LittleEndian.AppendUint16(code, uint16(number))
	default:
		return binary.LittleEndian.AppendUint32(code, uint32(number))
	}
}