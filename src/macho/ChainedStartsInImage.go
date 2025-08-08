package macho

const ChainedStartsInImageSize = 8

// ChainedStartsInImage describes which segments contain fixups.
type ChainedStartsInImage struct {
	SegCount      uint32
	SegInfoOffset [1]uint32
}