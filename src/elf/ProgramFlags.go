package elf

// ProgramFlags specifies the permissions for a segment.
type ProgramFlags int32

const (
	ProgramFlagsExecutable ProgramFlags = 0x1
	ProgramFlagsWritable   ProgramFlags = 0x2
	ProgramFlagsReadable   ProgramFlags = 0x4
)