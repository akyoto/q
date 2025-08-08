package macho

const UuidSize = 24

type Uuid struct {
	LoadCommand
	Length uint32
	Bytes  [16]byte
}