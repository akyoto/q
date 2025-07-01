package pe

const DOSHeaderSize = 64

// DOSHeader is at the beginning of each EXE file and nowadays just points
// to the PEHeader using an absolute file offset.
type DOSHeader struct {
	Magic          [4]byte
	_              [56]byte
	NTHeaderOffset uint32
}