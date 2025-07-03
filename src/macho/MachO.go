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
	PageZero   Segment64
	CodeHeader Segment64
	DataHeader Segment64
	UnixThread Thread
}

// Write writes the Mach-O format to the given writer.
func Write(writer io.WriteSeeker, b *build.Build, codeBytes []byte, dataBytes []byte) {
	x := exe.New(HeaderEnd, b.FileAlign(), b.MemoryAlign())
	x.InitSections(codeBytes, dataBytes)
	code := x.Sections[0]
	data := x.Sections[1]
	arch, microArch := Arch(b.Arch)
	entryPoint := BaseAddress + code.MemoryOffset

	m := &MachO{
		Header: Header{
			Magic:             0xFEEDFACF,
			Architecture:      arch,
			MicroArchitecture: microArch,
			Type:              TypeExecute,
			NumCommands:       4,
			SizeCommands:      SizeCommands,
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
			Name:         [16]byte{'_', '_', 'D', 'A', 'T', 'A'},
			Address:      uint64(BaseAddress + data.MemoryOffset),
			SizeInMemory: uint64(len(data.Bytes)),
			Offset:       uint64(data.FileOffset),
			SizeInFile:   uint64(len(data.Bytes)),
			NumSections:  0,
			Flag:         0,
			MaxProt:      ProtReadable,
			InitProt:     ProtReadable,
		},
		UnixThread: Thread{
			LoadCommand: LcUnixthread,
			Len:         ThreadSize,
			Type:        0x4,
			Data: [43]uint32{
				42,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
				uint32(entryPoint), 0,
				0, 0,
				0, 0,
				0, 0,
				0, 0,
			},
		},
	}

	binary.Write(writer, binary.LittleEndian, &m.Header)
	binary.Write(writer, binary.LittleEndian, &m.PageZero)
	binary.Write(writer, binary.LittleEndian, &m.CodeHeader)
	binary.Write(writer, binary.LittleEndian, &m.DataHeader)
	binary.Write(writer, binary.LittleEndian, &m.UnixThread)
	writer.Seek(int64(code.Padding), io.SeekCurrent)
	writer.Write(code.Bytes)
	writer.Seek(int64(data.Padding), io.SeekCurrent)
	writer.Write(data.Bytes)
}