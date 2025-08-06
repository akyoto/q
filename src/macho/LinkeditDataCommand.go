package macho

const LinkeditDataCommandSize = 16

// LinkeditDataCommand points to a data blob within the __LINKEDIT segment.
type LinkeditDataCommand struct {
	LoadCommand
	Length     uint32
	DataOffset uint32
	DataSize   uint32
}