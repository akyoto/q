package config

// Build describes the parameters for the "build" command.
type Build struct {
	Files         []string
	Arch          Arch
	OS            OS
	Filter        string
	Dry           bool
	FoldConstants bool
	RemoveCopies  bool
	ShowASM       bool
	ShowHeaders   bool
	ShowSSA       bool
}