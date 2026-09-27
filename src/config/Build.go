package config

// Build describes the parameters for the "build" command.
type Build struct {
	Files                []string
	Arch                 Arch
	OS                   OS
	Dry                  bool
	Fold                 bool
	RemoveCopies         bool
	Reorder              bool
	LintAssertionDensity bool
	LintBinaryOps        bool
	LintDeadCode         bool
	LintSlices           bool
	LintUnusedImports    bool
}