package elf

import "strings"

// AddSections adds section headers to the ELF file.
func (elf *ELF) AddSections() {
	elf.SectionHeaders = []SectionHeader{
		{
			Type: SectionTypeNULL,
		},
		{
			NameIndex:      int32(strings.Index(StringTable, ".text\000")),
			Type:           SectionTypePROGBITS,
			Flags:          SectionFlagsAllocate | SectionFlagsExecutable,
			VirtualAddress: elf.CodeHeader.VirtualAddress,
			Offset:         elf.CodeHeader.Offset,
			SizeInFile:     elf.CodeHeader.SizeInFile,
			Align:          elf.CodeHeader.Align,
		},
		{
			NameIndex:      int32(strings.Index(StringTable, ".rodata\000")),
			Type:           SectionTypePROGBITS,
			Flags:          SectionFlagsAllocate,
			VirtualAddress: elf.DataHeader.VirtualAddress,
			Offset:         elf.DataHeader.Offset,
			SizeInFile:     elf.DataHeader.SizeInFile,
			Align:          elf.DataHeader.Align,
		},
		{
			NameIndex:  int32(strings.Index(StringTable, ".shstrtab\000")),
			Type:       SectionTypeSTRTAB,
			Offset:     int64(StringTableStart),
			SizeInFile: int64(len(StringTable)),
			Align:      1,
		},
	}

	elf.SectionHeaderEntrySize = SectionHeaderSize
	elf.SectionHeaderEntryCount = int16(len(elf.SectionHeaders))
	elf.SectionHeaderOffset = int64(SectionHeaderStart)
	elf.SectionNameStringTableIndex = int16(len(elf.SectionHeaders) - 1)
}