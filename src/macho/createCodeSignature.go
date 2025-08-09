package macho

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math/bits"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/exe"
)

// createCodeSignature creates a signature for the binary.
func createCodeSignature(build *config.Build, code *exe.Section) []byte {
	superBlob := bytes.Buffer{}
	identifier := []byte("\000")
	pageSize := build.MemoryAlign()
	pageSizeExponent := bits.Len(uint(pageSize)) - 1
	numHashes := (len(code.Bytes) + pageSize - 1) / pageSize

	binary.Write(&superBlob, binary.BigEndian, &CodeDirectory{
		Magic:        CS_MAGIC_CODEDIRECTORY,
		Length:       uint32(CodeDirectorySize + len(identifier) + CS_SHA256_LEN),
		Version:      CS_SUPPORTSEXECSEG,
		HashOffset:   uint32(CodeDirectorySize + len(identifier)),
		HashSize:     CS_SHA256_LEN,
		HashType:     CS_HASHTYPE_SHA256,
		PageSize:     uint8(pageSizeExponent),
		Flags:        CS_ADHOC | CS_LINKER_SIGNED,
		NCodeSlots:   uint32(numHashes),
		CodeLimit:    uint32(len(code.Bytes)),
		IdentOffset:  CodeDirectorySize,
		ExecSegBase:  uint64(code.FileOffset),
		ExecSegLimit: uint64(len(code.Bytes)),
		ExecSegFlags: CS_EXECSEG_MAIN_BINARY,
	})

	// Identifier
	superBlob.Write(identifier)

	// Hashes
	for i := range numHashes {
		start := i * pageSize
		end := min(start+pageSize, len(code.Bytes))
		codeHash := sha256.Sum256(code.Bytes[start:end])
		superBlob.Write(codeHash[:])
	}

	return superBlob.Bytes()
}