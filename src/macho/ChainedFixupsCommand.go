package macho

const ChainedFixupsCommandSize = 16

// ChainedFixupsCommand points to a blob that starts with a ChainedFixupsHeader,
// enabling relocations in writable segments.
type ChainedFixupsCommand struct {
	LoadCommand
	Length     uint32
	DataOffset uint32
	DataSize   uint32
}