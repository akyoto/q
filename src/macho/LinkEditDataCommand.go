package macho

const LinkEditDataCommandSize = 16

// LinkEditDataCommand points to a data blob within the __LINKEDIT segment.
type LinkEditDataCommand struct {
	LoadCommand
	Length     uint32
	DataOffset uint32
	DataSize   uint32
}