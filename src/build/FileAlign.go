package build

// FileAlign returns the file alignment.
func (build *Build) FileAlign() int {
	return build.MemoryAlign()
}