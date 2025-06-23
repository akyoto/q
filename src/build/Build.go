package build

// Build describes the parameters for the "build" command.
type Build struct {
	Files       []string
	Arch        Arch
	OS          OS
	FileAlign   int
	MemoryAlign int
	Dry         bool
	ShowSSA     bool
}