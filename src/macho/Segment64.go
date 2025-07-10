package macho

const Segment64Size = 72

// Segment64 is a segment load command.
type Segment64 struct {
	LoadCommand
	Length       uint32
	Name         [16]byte
	Address      uint64
	SizeInMemory uint64
	Offset       uint64
	SizeInFile   uint64
	MaxProt      Prot
	InitProt     Prot
	NumSections  uint32
	Flag         SegmentFlags
}