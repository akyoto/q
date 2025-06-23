package build

// SetArch sets the architecture which also influences the default alignment.
func (build *Build) SetArch(arch Arch) {
	build.Arch = arch

	switch arch {
	case ARM:
		build.MemoryAlign = 0x4000
	default:
		build.MemoryAlign = 0x1000
	}

	build.FileAlign = build.MemoryAlign
}