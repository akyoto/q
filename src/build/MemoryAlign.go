package build

// MemoryAlign returns the memory alignment.
func (build *Build) MemoryAlign() int {
	switch build.Arch {
	case ARM:
		return 0x4000
	default:
		return 0x1000
	}
}