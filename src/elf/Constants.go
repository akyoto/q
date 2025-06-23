package elf

const (
	LittleEndian      = 1
	TypeExecutable    = 2
	TypeDynamic       = 3
	ArchitectureAMD64 = 0x3E
	ArchitectureARM64 = 0xB7
	ArchitectureRISCV = 0xF3
	HeaderEnd         = HeaderSize + ProgramHeaderSize*2
)