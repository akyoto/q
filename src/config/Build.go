package config

// Build describes the parameters for the "build" command.
type Build struct {
	Files         []string
	Arch          Arch
	OS            OS
	Filter        string
	Dry           bool
	Fold          bool
	Reorder       bool
	LintBinaryOps bool
}