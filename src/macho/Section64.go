package macho

const Section64Size = 80

// Section64 is one of multiple sections in a segment.
type Section64 struct {
	SectionName      [16]byte
	SegmentName      [16]byte
	Address          uint64
	Size             uint64
	Offset           uint32
	Align            uint32
	RelocationOffset uint32
	NumRelocations   uint32
	Flags            SectionFlags
	_                uint32
	_                uint32
	_                uint32
}