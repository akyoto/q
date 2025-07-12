package pe

import (
	"git.urbach.dev/cli/q/src/config"
)

const (
	IMAGE_FILE_MACHINE_AMD64   = 0x8664
	IMAGE_FILE_MACHINE_ARM64   = 0xAA64
	IMAGE_FILE_MACHINE_RISCV64 = 0x5064
)

// Arch returns the CPU architecture used in the PE header.
func Arch(arch config.Arch) uint16 {
	switch arch {
	case config.ARM:
		return IMAGE_FILE_MACHINE_ARM64
	case config.X86:
		return IMAGE_FILE_MACHINE_AMD64
	default:
		return 0
	}
}