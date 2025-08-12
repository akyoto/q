package pe

const (
	BaseAddress = 0x140000000
	NumSections = 3
	HeaderEnd   = DOSHeaderSize + NTHeaderSize + OptionalHeader64Size + SectionHeaderSize*NumSections
)