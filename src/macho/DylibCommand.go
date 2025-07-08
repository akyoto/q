package macho

const (
	DylibCommandSize = 24
	LibSystemString  = "/usr/lib/libSystem.B.dylib\000\000\000\000\000\000"
)

// DylibCommand is added for each shared library.
type DylibCommand struct {
	LoadCommand
	Length               uint32
	Name                 uint32
	TimeStamp            uint32
	CurrentVersion       uint32
	CompatibilityVersion uint32
}