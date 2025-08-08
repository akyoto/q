package macho

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math/bits"

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
		SegCount:      1,
		SegInfoOffset: [1]uint32{ChainedStartsInImageSize},
	})

	binary.Write(&buffer, binary.LittleEndian, &ChainedStartsInSegment{
		Size:          ChainedStartsInSegmentSize,
		PageSize:      uint16(build.MemoryAlign()),
		PageCount:     1,
		PageStarts:    [1]uint16{DYLD_CHAINED_PTR_START_NONE},
		PointerFormat: DYLD_CHAINED_PTR_64,
	})

	// Superblob
	superBlob := bytes.Buffer{}
	identifier := []byte("exe\000")
	pageSizeExponent := bits.Len(uint(build.MemoryAlign())) - 1

	binary.Write(&superBlob, binary.BigEndian, &CodeDirectory{
		Magic:        CS_MAGIC_CODEDIRECTORY,
		Length:       uint32(CodeDirectorySize + len(identifier) + CS_SHA256_LEN),
		Version:      CS_SUPPORTSEXECSEG,
		HashOffset:   uint32(CodeDirectorySize + len(identifier)),
		HashSize:     CS_SHA256_LEN,
		HashType:     CS_HASHTYPE_SHA256,
		PageSize:     uint8(pageSizeExponent),
		Flags:        CS_ADHOC | CS_LINKER_SIGNED,
		NCodeSlots:   1,
		CodeLimit:    uint32(len(code.Bytes)),
		IdentOffset:  CodeDirectorySize,
		ExecSegBase:  uint64(code.FileOffset),
		ExecSegLimit: uint64(len(code.Bytes)),
		ExecSegFlags: CS_EXECSEG_MAIN_BINARY,
	})

	paddedCode := append(code.Bytes, bytes.Repeat([]byte{0}, code.Padding)...)
	hasher := sha256.New()
	hasher.Write(paddedCode)
	codeHash := hasher.Sum(nil)

	superBlob.Write(identifier)
	superBlob.Write(codeHash)
	superBlobBytes := superBlob.Bytes()

	// Code signature
	offset, padding := exe.AlignPad(SuperBlobSize+BlobIndexSize, 8)

	binary.Write(&buffer, binary.BigEndian, &SuperBlob{
		Magic:  CS_MAGIC_EMBEDDED_SIGNATURE,
		Length: uint32(offset + len(superBlobBytes)),
		Count:  1,
	})

	binary.Write(&buffer, binary.BigEndian, &BlobIndex{
		Type:   CS_SLOT_CODEDIRECTORY,
		Offset: uint32(offset),
	})

	for range padding {
		buffer.WriteByte(0)
	}

	buffer.Write(superBlobBytes)
	return buffer.Bytes()
}