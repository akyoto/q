package macho

import "git.urbach.dev/cli/q/src/config"

// Arch returns the CPU architecture used in the Mach-O header.
func Arch(arch config.Arch) (CPU, uint32) {
	switch arch {
	case config.ARM:
		return CPU_ARM_64, CPU_SUBTYPE_ARM64_ALL
	case config.X86:
		return CPU_X86_64, CPU_SUBTYPE_X86_64_ALL
	default:
		return 0, 0
	}
}