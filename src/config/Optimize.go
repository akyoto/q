package config

// Optimize enables or disables the optimizer.
func (build *Build) Optimize(enabled bool) {
	build.FoldConstants = enabled
	build.RemoveCopies = enabled
}