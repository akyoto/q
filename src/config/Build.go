package config

// Build describes the parameters for the "build" command.
type Build struct {
	Files                []string
	Arch                 Arch
	OS                   OS
	Dry                  bool
	Fold                 bool
	Reorder              bool
	LintAssertionDensity bool
	LintBinaryOps        bool
	LintDeadCode         bool
	LintUnusedImports    bool
}