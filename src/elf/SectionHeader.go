package elf

// SectionHeaderSize is equal to the size of a section header in bytes.
const SectionHeaderSize = 64

// SectionHeader points to the data sections of our program.
type SectionHeader struct {
	NameIndex      int32
	Type           SectionType
	Flags          SectionFlags
	VirtualAddress int64
	Offset         int64
	SizeInFile     int64
	Link           int32
	Info           int32
	Align          int64
	EntrySize      int64
}