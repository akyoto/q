package config

// Arch is a CPU architecture.
type Arch uint8

const (
	UnknownArch Arch = iota
	ARM
	X86
)

// String returns the lowercase name of the architecture.
func (arch Arch) String() string {
	switch arch {
	case ARM:
		return "arm"
	case X86:
		return "x86"
	default:
		return "unknown"
	}
}