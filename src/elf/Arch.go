package elf

import "git.urbach.dev/cli/q/src/config"

// Arch converts the architecture variable to an ELF-specific constant.
func Arch(arch config.Arch) int16 {
	switch arch {
	case config.ARM:
		return ArchitectureARM64
	case config.X86:
		return ArchitectureAMD64
	default:
		return 0
	}
}