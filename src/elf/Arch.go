package elf

import "git.urbach.dev/cli/q/src/build"

// Arch converts the architecture variable to an ELF-specific constant.
func Arch(arch build.Arch) int16 {
	switch arch {
	case build.ARM:
		return ArchitectureARM64
	case build.X86:
		return ArchitectureAMD64
	default:
		return 0
	}
}