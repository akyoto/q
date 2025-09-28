package config

import (
	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/x86"
)

// CPU returns the correct CPU for the build.
func (build *Build) CPU() *cpu.CPU {
	switch build.Arch {
	case ARM:
		switch build.OS {
		case Linux:
			return &arm.LinuxCPU
		case Mac:
			return &arm.MacCPU
		case Windows:
			return &arm.WindowsCPU
		}
	case X86:
		switch build.OS {
		case Linux:
			return &x86.LinuxCPU
		case Mac:
			return &x86.MacCPU
		case Windows:
			return &x86.WindowsCPU
		}
	}

	return nil
}