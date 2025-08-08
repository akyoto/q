package macho

const ChainedStartsInImageSize = 4 + NumSegments*4

// ChainedStartsInImage describes which segments contain fixups.
type ChainedStartsInImage struct {
	SegCount      uint32
	SegInfoOffset [NumSegments]uint32
}