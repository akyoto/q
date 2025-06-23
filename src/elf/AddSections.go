package elf

import "bytes"

// AddSections adds section headers to the ELF file.
func (elf *ELF) AddSections() {
	elf.StringTable = []byte("\000.text\000.shstrtab\000")
	stringTableStart := elf.DataHeader.Offset + elf.DataHeader.SizeInFile
	sectionHeaderStart := stringTableStart + int64(len(elf.StringTable))

	elf.SectionHeaders = []SectionHeader{
		{
			Type: SectionTypeNULL,
		},
		{
			NameIndex:      int32(bytes.Index(elf.StringTable, []byte(".text\000"))),
			Type:           SectionTypePROGBITS,
			Flags:          SectionFlagsAllocate | SectionFlagsExecutable,
			VirtualAddress: elf.CodeHeader.VirtualAddress,
			Offset:         elf.CodeHeader.Offset,
			SizeInFile:     elf.CodeHeader.SizeInFile,
			Align:          elf.CodeHeader.Align,
		},
		{
			NameIndex:  int32(bytes.Index(elf.StringTable, []byte(".shstrtab\000"))),
			Type:       SectionTypeSTRTAB,
			Offset:     int64(stringTableStart),
			SizeInFile: int64(len(elf.StringTable)),
			Align:      1,
		},
	}

	elf.SectionHeaderEntrySize = SectionHeaderSize
	elf.SectionHeaderEntryCount = int16(len(elf.SectionHeaders))
	elf.SectionHeaderOffset = int64(sectionHeaderStart)
	elf.SectionNameStringTableIndex = 2
}