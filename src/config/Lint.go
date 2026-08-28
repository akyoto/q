package config

// Lint enables or disables the linter.
func (build *Build) Lint(enabled bool) {
	build.LintAssertionDensity = enabled
	build.LintBinaryOps = enabled
	build.LintDeadCode = enabled
}