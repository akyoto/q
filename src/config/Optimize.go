package config

// Optimize enables or disables the optimizer.
func (build *Build) Optimize(enabled bool) {
	build.Fold = enabled
	build.Reorder = enabled
}