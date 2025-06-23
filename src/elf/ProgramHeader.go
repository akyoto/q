package elf

// ProgramHeaderSize is equal to the size of a program header in bytes.
const ProgramHeaderSize = 56

// ProgramHeader points to the executable part of our program.
type ProgramHeader struct {
	Type            ProgramType
	Flags           ProgramFlags
	Offset          int64
	VirtualAddress  int64
	PhysicalAddress int64
	SizeInFile      int64
	SizeInMemory    int64
	Align           int64
}