package macho

const MainSize = 24

// Main is the structure of the LC_MAIN load command.
type Main struct {
	LoadCommand
	Length          uint32
	EntryFileOffset uint64
	StackSize       uint64
}