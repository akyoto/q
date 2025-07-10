package macho

import (
	"encoding/binary"
	"io"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/exe"
)

// MachO is the executable format used on MacOS.
type MachO struct {
	Header
	PageZero        Segment64
	CodeSegment     Segment64
	CodeSection     Section64
	DataSegment     Segment64
	MainHeader      Main
	InfoHeader      DyldInfoCommand
	LinkerHeader    DylinkerCommand
	LibSystemHeader DylibCommand
}

// Write writes the Mach-O format to the given writer.
func Write(writer io.WriteSeeker, b *build.Build, codeBytes []byte, dataBytes []byte) {
	x := exe.New(HeaderEnd, b.FileAlign(), b.MemoryAlign(), b.Congruent(), codeBytes, dataBytes)
	code := x.Sections[0]
	data := x.Sections[1]
	arch, microArch := Arch(b.Arch)

	m := &MachO{
		Header: Header{
			Magic:             0xFEEDFACF,
			Architecture:      arch,
			MicroArchitecture: microArch,
			Type:              TypeExecute,
			NumCommands:       NumCommands,
			SizeCommands:      uint32(SizeCommands),
			Flags:             FlagNoUndefs | FlagDyldLink | FlagTwoLevel | FlagPIE | FlagNoHeapExecution,
		},
		PageZero: Segment64{
			LoadCommand:  LcSegment64,
			Length:       72,
			Name:         [16]byte{'_', '_', 'P', 'A', 'G', 'E', 'Z', 'E', 'R', 'O'},
			Address:      0,
			SizeInMemory: uint64(BaseAddress),
			Offset:       0,
			SizeInFile:   0,
			NumSections:  0,
			Flag:         0,
			MaxProt:      0,
			InitProt:     0,
		},
		CodeSegment: Segment64{
			LoadCommand:  LcSegment64,
			Length:       Segment64Size + Section64Size,
			Name:         [16]byte{'_', '_', 'T', 'E', 'X', 'T'},
			Address:      uint64(BaseAddress),
			SizeInMemory: uint64(code.MemoryOffset + len(code.Bytes)),
			Offset:       0,
			SizeInFile:   uint64(code.FileOffset + len(code.Bytes)),
			NumSections:  1,
			Flag:         0,
			MaxProt:      ProtReadable | ProtExecutable,
			InitProt:     ProtReadable | ProtExecutable,
		},
		CodeSection: Section64{
			SectionName: [16]byte{'_', '_', 't', 'e', 'x', 't'},
			SegmentName: [16]byte{'_', '_', 'T', 'E', 'X', 'T'},
			Address:     uint64(BaseAddress + code.MemoryOffset),
			Size:        uint64(len(code.Bytes)),
			Offset:      uint32(code.FileOffset),
			Align:       6,
			Flags:       SectionPureInstructions,
		},
		DataSegment: Segment64{
			LoadCommand:  LcSegment64,
			Length:       Segment64Size,
			Name:         [16]byte{'_', '_', 'D', 'A', 'T', 'A', '_', 'C', 'O', 'N', 'S', 'T'},
			Address:      uint64(BaseAddress + data.MemoryOffset),
			SizeInMemory: uint64(len(data.Bytes)),
			Offset:       uint64(data.FileOffset),
			SizeInFile:   uint64(len(data.Bytes)),
			NumSections:  0,
			Flag:         SegmentReadOnly,
			MaxProt:      ProtReadable,
			InitProt:     ProtReadable,
		},
		MainHeader: Main{
			LoadCommand:     LcMain,
			Length:          MainSize,
			EntryFileOffset: uint64(code.MemoryOffset),
			StackSize:       0,
		},
		InfoHeader: DyldInfoCommand{
			LoadCommand: LcDyldInfoOnly,
			Length:      DyldInfoCommandSize,
		},
		LinkerHeader: DylinkerCommand{
			LoadCommand: LcLoadDylinker,
			Length:      uint32(DylinkerCommandSize + len(LinkerString)),
			Name:        DylinkerCommandSize,
		},
		LibSystemHeader: DylibCommand{
			LoadCommand: LcLoadDylib,
			Length:      uint32(DylibCommandSize + len(LibSystemString)),
			Name:        DylibCommandSize,
		},
	}

	binary.Write(writer, binary.LittleEndian, &m.Header)
	binary.Write(writer, binary.LittleEndian, &m.PageZero)
	binary.Write(writer, binary.LittleEndian, &m.CodeSegment)
	binary.Write(writer, binary.LittleEndian, &m.CodeSection)
	binary.Write(writer, binary.LittleEndian, &m.DataSegment)
	binary.Write(writer, binary.LittleEndian, &m.MainHeader)
	binary.Write(writer, binary.LittleEndian, &m.InfoHeader)
	binary.Write(writer, binary.LittleEndian, &m.LinkerHeader)
	writer.Write([]byte(LinkerString))
	binary.Write(writer, binary.LittleEndian, &m.LibSystemHeader)
	writer.Write([]byte(LibSystemString))
	writer.Seek(int64(code.Padding), io.SeekCurrent)
	writer.Write(code.Bytes)
	writer.Seek(int64(data.Padding), io.SeekCurrent)
	writer.Write(data.Bytes)
}