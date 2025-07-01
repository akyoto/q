package pe

const (
	BaseAddress = 0x400000
	NumSections = 3
	HeaderEnd   = DOSHeaderSize + NTHeaderSize + OptionalHeader64Size + SectionHeaderSize*NumSections
)