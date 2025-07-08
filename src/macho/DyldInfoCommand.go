package macho

const DyldInfoCommandSize = 48

// DyldInfoCommand contains the file offsets and sizes that dyld needs to load the image.
type DyldInfoCommand struct {
	LoadCommand
	Length         uint32
	RebaseOffset   uint32
	RebaseSize     uint32
	BindOffset     uint32
	BindSize       uint32
	WeakBindOffset uint32
	WeakBindSize   uint32
	LazyBindOffset uint32
	LazyBindSize   uint32
	ExportOffset   uint32
	ExportSize     uint32
}