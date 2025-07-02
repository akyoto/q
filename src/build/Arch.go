package build

// Arch is a CPU architecture.
type Arch uint8

const (
	UnknownArch Arch = iota
	ARM
	X86
)

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