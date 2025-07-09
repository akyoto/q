package build

// MemoryAlign returns the memory alignment.
func (build *Build) MemoryAlign() int {
	switch build.Arch {
	case ARM:
		switch build.OS {
		case Linux:
			return 0x10000
		case Mac:
			return 0x4000
		default:
			return 0x1000
		}

	default:
		return 0x1000
	}
}