package macho

const (
	DylinkerCommandSize = 12
	LinkerString        = "/usr/lib/dyld\000\000\000\000\000\000\000"
)

// DylinkerCommand is needed if the program uses a dynamic linker.
type DylinkerCommand struct {
	LoadCommand
	Length uint32
	Name   uint32
}