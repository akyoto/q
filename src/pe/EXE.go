package pe

import (
	"bytes"
	"encoding/binary"
	"io"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/dll"
	"git.urbach.dev/cli/q/src/exe"
)

// EXE is the portable executable format used on Windows.
type EXE struct {
	DOSHeader
	NTHeader
	OptionalHeader64
	Sections []SectionHeader
}

// Write writes the EXE file to the given writer.
func Write(writer io.WriteSeeker, b *build.Build, codeBytes []byte, dataBytes []byte, libs dll.List) {
	x := exe.New(HeaderEnd, b.FileAlign(), b.MemoryAlign(), b.Congruent(), codeBytes, dataBytes, nil)
	code := x.Sections[0]
	data := x.Sections[1]
	imports := x.Sections[2]
	subSystem := IMAGE_SUBSYSTEM_WINDOWS_CUI
	arch := Arch(b.Arch)

	importsData, dllData, dllImports, dllDataStart := importLibraries(libs, imports.MemoryOffset)
	buffer := bytes.Buffer{}
	binary.Write(&buffer, binary.LittleEndian, &importsData)
	binary.Write(&buffer, binary.LittleEndian, &dllData)
	binary.Write(&buffer, binary.LittleEndian, &dllImports)
	imports.Bytes = buffer.Bytes()
	importDirectoryStart := dllDataStart + len(dllData)
	importDirectorySize := DLLImportSize * len(dllImports)
	imageSize := exe.Align(imports.MemoryOffset+len(imports.Bytes), b.MemoryAlign())

	if libs.Contains("user32") {
		subSystem = IMAGE_SUBSYSTEM_WINDOWS_GUI
	}

	pe := &EXE{
		DOSHeader: DOSHeader{
			Magic:          [4]byte{'M', 'Z', 0, 0},
			NTHeaderOffset: DOSHeaderSize,
		},
		NTHeader: NTHeader{
			Signature:            [4]byte{'P', 'E', 0, 0},
			Machine:              arch,
			NumberOfSections:     uint16(NumSections),
			TimeDateStamp:        0,
			PointerToSymbolTable: 0,
			NumberOfSymbols:      0,
			SizeOfOptionalHeader: OptionalHeader64Size,
			Characteristics:      IMAGE_FILE_EXECUTABLE_IMAGE | IMAGE_FILE_LARGE_ADDRESS_AWARE,
		},
		OptionalHeader64: OptionalHeader64{
			Magic:                       0x020B, // PE32+ / 64-bit executable
			MajorLinkerVersion:          0x0E,
			MinorLinkerVersion:          0x16,
			SizeOfCode:                  uint32(len(code.Bytes)),
			SizeOfInitializedData:       0,
			SizeOfUninitializedData:     0,
			AddressOfEntryPoint:         uint32(code.MemoryOffset),
			BaseOfCode:                  uint32(code.MemoryOffset),
			ImageBase:                   BaseAddress,
			SectionAlignment:            uint32(b.MemoryAlign()), // power of 2, must be greater than or equal to FileAlignment
			FileAlignment:               uint32(b.FileAlign()),   // power of 2
			MajorOperatingSystemVersion: 0x06,
			MinorOperatingSystemVersion: 0,
			MajorImageVersion:           0,
			MinorImageVersion:           0,
			MajorSubsystemVersion:       0x06,
			MinorSubsystemVersion:       0,
			Win32VersionValue:           0,
			SizeOfImage:                 uint32(imageSize),       // a multiple of SectionAlignment
			SizeOfHeaders:               uint32(code.FileOffset), // section bodies begin here
			CheckSum:                    0,
			Subsystem:                   subSystem,
			DllCharacteristics:          IMAGE_DLLCHARACTERISTICS_HIGH_ENTROPY_VA | IMAGE_DLLCHARACTERISTICS_DYNAMIC_BASE | IMAGE_DLLCHARACTERISTICS_NX_COMPAT | IMAGE_DLLCHARACTERISTICS_TERMINAL_SERVER_AWARE,
			SizeOfStackReserve:          0x100000,
			SizeOfStackCommit:           0x1000,
			SizeOfHeapReserve:           0x100000,
			SizeOfHeapCommit:            0x1000,
			LoaderFlags:                 0,
			NumberOfRvaAndSizes:         16,
			DataDirectory: [16]DataDirectory{
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: uint32(importDirectoryStart), Size: uint32(importDirectorySize)}, // RVA of the imported function table
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: uint32(imports.MemoryOffset), Size: uint32(len(importsData) * 8)}, // RVA of the import address table
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: 0, Size: 0},
				{VirtualAddress: 0, Size: 0},
			},
		},
		Sections: []SectionHeader{
			{
				Name:            [8]byte{'.', 't', 'e', 'x', 't'},
				VirtualSize:     uint32(len(code.Bytes)),
				VirtualAddress:  uint32(code.MemoryOffset),
				RawSize:         uint32(len(code.Bytes)), // must be a multiple of FileAlignment
				RawAddress:      uint32(code.FileOffset), // must be a multiple of FileAlignment
				Characteristics: IMAGE_SCN_CNT_CODE | IMAGE_SCN_MEM_EXECUTE | IMAGE_SCN_MEM_READ,
			},
			{
				Name:            [8]byte{'.', 'r', 'd', 'a', 't', 'a'},
				VirtualSize:     uint32(len(data.Bytes)),
				VirtualAddress:  uint32(data.MemoryOffset),
				RawSize:         uint32(len(data.Bytes)),
				RawAddress:      uint32(data.FileOffset),
				Characteristics: IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ,
			},
			{
				Name:            [8]byte{'.', 'i', 'd', 'a', 't', 'a'},
				VirtualSize:     uint32(len(imports.Bytes)),
				VirtualAddress:  uint32(imports.MemoryOffset),
				RawSize:         uint32(len(imports.Bytes)),
				RawAddress:      uint32(imports.FileOffset),
				Characteristics: IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ,
			},
		},
	}

	binary.Write(writer, binary.LittleEndian, &pe.DOSHeader)
	binary.Write(writer, binary.LittleEndian, &pe.NTHeader)
	binary.Write(writer, binary.LittleEndian, &pe.OptionalHeader64)
	binary.Write(writer, binary.LittleEndian, &pe.Sections)
	writer.Seek(int64(code.Padding), io.SeekCurrent)
	writer.Write(code.Bytes)
	writer.Seek(int64(data.Padding), io.SeekCurrent)
	writer.Write(data.Bytes)
	writer.Seek(int64(imports.Padding), io.SeekCurrent)
	writer.Write(imports.Bytes)
}