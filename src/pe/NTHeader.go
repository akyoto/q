package pe

const NTHeaderSize = 24

type NTHeader struct {
	Signature            [4]byte
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      Characteristics
}