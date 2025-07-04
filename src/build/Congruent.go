package build

// Congruent returns true if the platform needs to force `virtual address % alignment == file offset % alignment`.
func (build *Build) Congruent() bool {
	return build.OS == Linux
}