package macho

import (
	"bytes"
	"encoding/binary"
	"io"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/exe"
)

// MachO is the executable format used on MacOS.
type MachO struct {
	*exe.Executable
	Header
	PageZero       Segment64
	CodeSegment    Segment64
	CodeSection    Section64
	DataSegment    Segment64
	ImportsSegment Segment64
	Uuid           Uuid
	Main           Main
	BuildVersion   BuildVersion
	Linker         DylinkerCommand
	LibSystem      DylibCommand
	ChainedFixups  LinkeditDataCommand
	CodeSignature  LinkeditDataCommand
}

// Write writes the Mach-O format to the given writer.
func Write(writer io.WriteSeeker, build *config.Build, codeBytes []byte, dataBytes []byte) {
	x := exe.New(HeaderEnd, build.FileAlign(), build.MemoryAlign(), build.Congruent(), codeBytes, dataBytes, createLinkeditSegment())
	code := x.Sections[0]
	data := x.Sections[1]
	imports := x.Sections[2]
	arch, microArch := Arch(build.Arch)
	chainedFixupsSize := ChainedFixupsHeaderSize + ChainedStartsInImageSize + CodeSignaturePadding
	rawFileSize := imports.FileOffset + len(imports.Bytes)
	identifier := []byte("\000")
	codeSignature := NewCodeSignature(rawFileSize, identifier, code)

	m := &MachO{
		Executable: x,
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
			MaxProt:      ProtReadable | ProtWritable, // dyld complains if it's not writable
			InitProt:     ProtReadable | ProtWritable,
		},
		ImportsSegment: Segment64{
			LoadCommand:  LcSegment64,
			Length:       Segment64Size,
			Name:         [16]byte{'_', '_', 'L', 'I', 'N', 'K', 'E', 'D', 'I', 'T'},
			Address:      uint64(BaseAddress + imports.MemoryOffset),
			SizeInMemory: uint64(len(imports.Bytes) + int(codeSignature.size())),
			Offset:       uint64(imports.FileOffset),
			SizeInFile:   uint64(len(imports.Bytes) + int(codeSignature.size())),
			NumSections:  0,
			Flag:         0,
			MaxProt:      ProtReadable,
			InitProt:     ProtReadable,
		},
		Uuid: Uuid{
			LoadCommand: LcUuid,
			Length:      UuidSize,
			Bytes:       [16]byte{},
		},
		Main: Main{
			LoadCommand:     LcMain,
			Length:          MainSize,
			EntryFileOffset: uint64(code.MemoryOffset),
			StackSize:       0,
		},
		BuildVersion: BuildVersion{
			LoadCommand: LcBuildVersion,
			Length:      BuildVersionSize,
			Platform:    PlatformMacOS,
			MinOS:       Version(12, 0, 0),
			Sdk:         Version(12, 0, 0),
			NumTools:    0,
		},
		Linker: DylinkerCommand{
			LoadCommand: LcLoadDylinker,
			Length:      uint32(DylinkerCommandSize + len(LinkerString)),
			Name:        DylinkerCommandSize,
		},
		LibSystem: DylibCommand{
			LoadCommand: LcLoadDylib,
			Length:      uint32(DylibCommandSize + len(LibSystemString)),
			Name:        DylibCommandSize,
		},
		ChainedFixups: LinkeditDataCommand{
			LoadCommand: LcDyldChainedFixups,
			Length:      LinkeditDataCommandSize,
			DataOffset:  uint32(imports.FileOffset),
			DataSize:    uint32(chainedFixupsSize),
		},
		CodeSignature: LinkeditDataCommand{
			LoadCommand: LcCodeSignature,
			Length:      LinkeditDataCommandSize,
			DataOffset:  uint32(imports.FileOffset + chainedFixupsSize),
			DataSize:    codeSignature.size(),
		},
	}

	buffer := bytes.Buffer{}
	m.WriteRaw(&buffer, false)
	contents := buffer.Bytes()
	m.WriteRaw(writer, true)
	codeSignature.write(writer, contents, identifier)
}

// WriteRaw writes the raw file contents without the code signature.
func (m *MachO) WriteRaw(writer io.Writer, seek bool) {
	binary.Write(writer, binary.LittleEndian, &m.Header)
	binary.Write(writer, binary.LittleEndian, &m.PageZero)
	binary.Write(writer, binary.LittleEndian, &m.CodeSegment)
	binary.Write(writer, binary.LittleEndian, &m.CodeSection)
	binary.Write(writer, binary.LittleEndian, &m.DataSegment)
	binary.Write(writer, binary.LittleEndian, &m.ImportsSegment)
	binary.Write(writer, binary.LittleEndian, &m.Uuid)
	binary.Write(writer, binary.LittleEndian, &m.Main)
	binary.Write(writer, binary.LittleEndian, &m.BuildVersion)
	binary.Write(writer, binary.LittleEndian, &m.Linker)
	writer.Write([]byte(LinkerString))
	binary.Write(writer, binary.LittleEndian, &m.LibSystem)
	writer.Write([]byte(LibSystemString))
	binary.Write(writer, binary.LittleEndian, &m.ChainedFixups)
	binary.Write(writer, binary.LittleEndian, &m.CodeSignature)

	for _, section := range m.Executable.Sections {
		if seek {
			writer.(io.WriteSeeker).Seek(int64(section.Padding), io.SeekCurrent)
		} else {
			writer.Write(bytes.Repeat([]byte{0}, section.Padding))
		}

		writer.Write(section.Bytes)
	}
}