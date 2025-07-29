package config

// SetVerbose activates or deactivates verbose output.
func (build *Build) SetVerbose(verbose bool) {
	build.ShowASM = verbose
	build.ShowHeaders = verbose
	build.ShowIR = verbose
}