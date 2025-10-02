package macho

import (
	"bytes"
	"encoding/binary"
)

const (
	codeSignaturePadding = 12
	chainedFixupsSize    = ChainedFixupsHeaderSize + ChainedStartsInImageSize + codeSignaturePadding
)

// createLinkeditSegment creates the contents of the __LINKEDIT segment.
func createLinkeditSegment() []byte {
	buffer := bytes.Buffer{}

	// Chained fixups
	binary.Write(&buffer, binary.LittleEndian, &ChainedFixupsHeader{
		StartsOffset:  ChainedFixupsHeaderSize,
		ImportsFormat: DYLD_CHAINED_IMPORT,
	})

	binary.Write(&buffer, binary.LittleEndian, &ChainedStartsInImage{
		SegCount:      NumSegments,
		SegInfoOffset: [NumSegments]uint32{},
	})

	// Make sure the code signature that follows is 16-byte aligned
	for range codeSignaturePadding {
		buffer.WriteByte(0)
	}

	return buffer.Bytes()
}