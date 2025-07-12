package config

// FileAlign returns the file alignment.
func (build *Build) FileAlign() int {
	switch build.OS {
	case Linux:
		return 0x40
	case Windows:
		return 0x200
	default:
		return build.MemoryAlign()
	}
}