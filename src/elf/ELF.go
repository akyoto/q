package elf

import (
	"encoding/binary"
	"io"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/exe"
)

// ELF represents an ELF file.
type ELF struct {
	Header
	CodeHeader     ProgramHeader
	DataHeader     ProgramHeader
	SectionHeaders []SectionHeader
}

// Write writes the ELF64 format to the given writer.
func Write(writer io.WriteSeeker, build *config.Build, codeBytes []byte, dataBytes []byte) {
	x := exe.New(HeaderEnd, build.FileAlign(), build.MemoryAlign(), build.Congruent(), false, codeBytes, dataBytes)
	code := x.Sections[0]
	data := x.Sections[1]

	elf := &ELF{
		Header: Header{
			Magic:                       [4]byte{0x7F, 'E', 'L', 'F'},
			Class:                       2,
			Endianness:                  LittleEndian,
			Version:                     1,
			OSABI:                       0,
			ABIVersion:                  0,
			Type:                        TypeDynamic,
			Architecture:                Arch(build.Arch),
			FileVersion:                 1,
			EntryPointInMemory:          int64(code.MemoryOffset),
			ProgramHeaderOffset:         HeaderSize,
			SectionHeaderOffset:         0,
			Flags:                       0,
			Size:                        HeaderSize,
			ProgramHeaderEntrySize:      ProgramHeaderSize,
			ProgramHeaderEntryCount:     2,
			SectionHeaderEntrySize:      0,
			SectionHeaderEntryCount:     0,
			SectionNameStringTableIndex: 0,
		},
		CodeHeader: ProgramHeader{
			Type:           ProgramTypeLOAD,
			Flags:          ProgramFlagsExecutable | ProgramFlagsReadable,
			Offset:         int64(code.FileOffset),
			VirtualAddress: int64(code.MemoryOffset),
			SizeInFile:     int64(len(code.Bytes)),
			SizeInMemory:   int64(len(code.Bytes)),
			Align:          int64(build.MemoryAlign()),
		},
		DataHeader: ProgramHeader{
			Type:           ProgramTypeLOAD,
			Flags:          ProgramFlagsReadable,
			Offset:         int64(data.FileOffset),
			VirtualAddress: int64(data.MemoryOffset),
			SizeInFile:     int64(len(data.Bytes)),
			SizeInMemory:   int64(len(data.Bytes)),
			Align:          int64(build.MemoryAlign()),
		},
	}

	elf.AddSections()
	binary.Write(writer, binary.LittleEndian, &elf.Header)
	binary.Write(writer, binary.LittleEndian, &elf.CodeHeader)
	binary.Write(writer, binary.LittleEndian, &elf.DataHeader)
	writer.Write([]byte(StringTable))
	binary.Write(writer, binary.LittleEndian, &elf.SectionHeaders)
	writer.Seek(int64(code.Padding), io.SeekCurrent)
	writer.Write(code.Bytes)
	writer.Seek(int64(data.Padding), io.SeekCurrent)
	writer.Write(data.Bytes)
}