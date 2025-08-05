package macho

// ChainedStartsInImage describes which segments contain fixups.
type ChainedStartsInImage struct {
	SegCount      uint32
	SegInfoOffset [1]uint32
}