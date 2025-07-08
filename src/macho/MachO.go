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
	CodeHeader      Segment64
	DataHeader      Segment64
	ImportsHeader   Segment64
	MainHeader      Main
	InfoHeader      DyldInfoCommand
	LinkerHeader    DylinkerCommand
	LibSystemHeader DylibCommand
}

// Write writes the Mach-O format to the given writer.
func Write(writer io.WriteSeeker, b *build.Build, codeBytes []byte, dataBytes []byte) {
	x := exe.New(HeaderEnd, b.FileAlign(), b.MemoryAlign(), b.Congruent(), codeBytes, dataBytes, nil)
	code := x.Sections[0]
	data := x.Sections[1]
	imports := x.Sections[2]
	arch, microArch := Arch(b.Arch)

	m := &MachO{
		Header: Header{
			Magic:             0xFEEDFACF,
			Architecture:      arch,
			MicroArchitecture: microArch,
			Type:              TypeExecute,
			NumCommands:       8,
			SizeCommands:      uint32(SizeCommands),
			Flags:             FlagNoUndefs | FlagPIE | FlagNoHeapExecution,
			Reserved:          0,
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
		CodeHeader: Segment64{
			LoadCommand:  LcSegment64,
			Length:       Segment64Size,
			Name:         [16]byte{'_', '_', 'T', 'E', 'X', 'T'},
			Address:      uint64(BaseAddress),
			SizeInMemory: uint64(code.MemoryOffset + len(code.Bytes)),
			Offset:       0,
			SizeInFile:   uint64(code.FileOffset + len(code.Bytes)),
			NumSections:  0,
			Flag:         0,
			MaxProt:      ProtReadable | ProtExecutable,
			InitProt:     ProtReadable | ProtExecutable,
		},
		DataHeader: Segment64{
			LoadCommand:  LcSegment64,
			Length:       Segment64Size,
			Name:         [16]byte{'_', '_', 'D', 'A', 'T', 'A', '_', 'C', 'O', 'N', 'S', 'T'},
			Address:      uint64(BaseAddress + data.MemoryOffset),
			SizeInMemory: uint64(len(data.Bytes)),
			Offset:       uint64(data.FileOffset),
			SizeInFile:   uint64(len(data.Bytes)),
			NumSections:  0,
			Flag:         0,
			MaxProt:      ProtReadable,
			InitProt:     ProtReadable,
		},
		ImportsHeader: Segment64{
			LoadCommand:  LcSegment64,
			Length:       Segment64Size,
			Name:         [16]byte{'_', '_', 'L', 'I', 'N', 'K', 'E', 'D', 'I', 'T'},
			Address:      uint64(BaseAddress + imports.MemoryOffset),
			SizeInMemory: uint64(len(imports.Bytes)),
			Offset:       uint64(imports.FileOffset),
			SizeInFile:   uint64(len(imports.Bytes)),
			NumSections:  0,
			Flag:         0,
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
			LoadCommand:  LcDyldInfoOnly,
			Length:       DyldInfoCommandSize,
			RebaseOffset: uint32(imports.FileOffset),
			RebaseSize:   uint32(len(imports.Bytes)),
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
	binary.Write(writer, binary.LittleEndian, &m.CodeHeader)
	binary.Write(writer, binary.LittleEndian, &m.DataHeader)
	binary.Write(writer, binary.LittleEndian, &m.ImportsHeader)
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
	writer.Seek(int64(imports.Padding), io.SeekCurrent)
	writer.Write(imports.Bytes)
}