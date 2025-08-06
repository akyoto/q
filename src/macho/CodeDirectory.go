package macho

const CodeDirectorySize = 88

// CodeDirectory points to the hashes of the machine code.
type CodeDirectory struct {
	Magic         uint32
	Length        uint32
	Version       uint32
	Flags         uint32
	HashOffset    uint32
	IdentOffset   uint32
	NSpecialSlots uint32
	NCodeSlots    uint32
	CodeLimit     uint32
	HashSize      uint8
	HashType      uint8
	Platform      uint8
	PageSize      uint8
	Spare2        uint32

	// Version 0x20100
	ScatterOffset uint32

	// Version 0x20200
	TeamOffset uint32

	// Version 0x20300
	Spare3      uint32
	CodeLimit64 uint64

	// Version 0x20400
	ExecSegBase  uint64
	ExecSegLimit uint64
	ExecSegFlags uint64
}