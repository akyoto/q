package elf

const (
	LittleEndian       = 1
	TypeExecutable     = 2
	TypeDynamic        = 3
	ArchitectureAMD64  = 0x3E
	ArchitectureARM64  = 0xB7
	ArchitectureRISCV  = 0xF3
	StringTable        = "\000.text\000.rodata\000.shstrtab\000"
	StringTableStart   = ProgramHeaderEnd
	SectionHeaderStart = StringTableStart + len(StringTable)
	ProgramHeaderEnd   = HeaderSize + ProgramHeaderSize*2
	HeaderEnd          = ProgramHeaderEnd + len(StringTable) + 4*SectionHeaderSize
)