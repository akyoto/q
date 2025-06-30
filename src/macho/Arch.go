package macho

import "git.urbach.dev/cli/q/src/build"

// Arch returns the CPU architecture used in the Mach-O header.
func Arch(arch build.Arch) (CPU, uint32) {
	switch arch {
	case build.ARM:
		return CPU_ARM_64, CPU_SUBTYPE_ARM64_ALL | 0x80000000
	case build.X86:
		return CPU_X86_64, CPU_SUBTYPE_X86_64_ALL | 0x80000000
	default:
		return 0, 0
	}
}