package macho

import (
	"bytes"
	"encoding/binary"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/exe"
)

// createLinkeditSegment creates the contents of the __LINKEDIT segment.
func createLinkeditSegment(build *config.Build, code *exe.Section) []byte {
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
	for range CodeSignaturePadding {
		buffer.WriteByte(0)
	}

	// Code signature
	// signature := createCodeSignature(build, code)
	// offset, padding := exe.AlignPad(SuperBlobSize+BlobIndexSize, 8)

	// binary.Write(&buffer, binary.BigEndian, &SuperBlob{
	// 	Magic:  CS_MAGIC_EMBEDDED_SIGNATURE,
	// 	Length: uint32(offset + len(signature)),
	// 	Count:  1,
	// })

	// binary.Write(&buffer, binary.BigEndian, &BlobIndex{
	// 	Type:   CS_SLOT_CODEDIRECTORY,
	// 	Offset: uint32(offset),
	// })

	// for range padding {
	// 	buffer.WriteByte(0)
	// }

	// buffer.Write(signature)
	return buffer.Bytes()
}