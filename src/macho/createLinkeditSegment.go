package macho

import (
	"bytes"
	"encoding/binary"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/exe"
)

const (
	// Code signing flags
	CS_ADHOC                    = 0x00000002
	CS_HARD                     = 0x00000100
	CS_KILL                     = 0x00000200
	CS_CHECK_EXPIRATION         = 0x00000400
	CS_RESTRICT                 = 0x00000800
	CS_ENFORCEMENT              = 0x00001000
	CS_REQUIRE_LV               = 0x00002000
	CS_RUNTIME                  = 0x00010000
	CS_LINKER_SIGNED            = 0x00020000
	CS_ALLOWED_MACHO            = CS_ADHOC | CS_HARD | CS_KILL | CS_CHECK_EXPIRATION | CS_RESTRICT | CS_ENFORCEMENT | CS_REQUIRE_LV | CS_RUNTIME | CS_LINKER_SIGNED
	CS_HASHTYPE_SHA256          = 2
	CS_SHA256_LEN               = 32
	CS_SLOT_CODEDIRECTORY       = 0
	CS_MAGIC_CODEDIRECTORY      = 0xFADE0C02
	CS_MAGIC_EMBEDDED_SIGNATURE = 0xFADE0CC0
	CS_SUPPORTSEXECSEG          = 0x20400
	CS_EXECSEG_MAIN_BINARY      = 0x1

	// Dyld
	DYLD_CHAINED_IMPORT         = 1
	DYLD_CHAINED_PTR_64         = 2
	DYLD_CHAINED_PTR_START_NONE = 0xFFFF

	// Custom
	CodeSignaturePadding = 4
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
	signature := createCodeSignature(build, code)
	offset, padding := exe.AlignPad(SuperBlobSize+BlobIndexSize, 8)

	binary.Write(&buffer, binary.BigEndian, &SuperBlob{
		Magic:  CS_MAGIC_EMBEDDED_SIGNATURE,
		Length: uint32(offset + len(signature)),
		Count:  1,
	})

	binary.Write(&buffer, binary.BigEndian, &BlobIndex{
		Type:   CS_SLOT_CODEDIRECTORY,
		Offset: uint32(offset),
	})

	for range padding {
		buffer.WriteByte(0)
	}

	buffer.Write(signature)
	return buffer.Bytes()
}