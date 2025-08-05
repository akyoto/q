package macho

const ChainedStartsInSegmentSize = 24

// ChainedStartsInSegment describes how pointer fixups are laid out in a segment.
type ChainedStartsInSegment struct {
	Size            uint32
	PageSize        uint16
	PointerFormat   uint16
	SegmentOffset   uint64
	MaxValidPointer uint32
	PageCount       uint16
	PageStarts      [1]uint16
}