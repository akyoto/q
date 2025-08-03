package codegen

import (
	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/x86"
)

// selectCPU returns the correct CPU for the build.
func selectCPU(build *config.Build) *cpu.CPU {
	switch build.Arch {
	case config.ARM:
		switch build.OS {
		case config.Linux:
			return &arm.LinuxCPU
		case config.Mac:
			return &arm.MacCPU
		case config.Windows:
			return &arm.WindowsCPU
		}
	case config.X86:
		switch build.OS {
		case config.Linux:
			return &x86.LinuxCPU
		case config.Mac:
			return &x86.MacCPU
		case config.Windows:
			return &x86.WindowsCPU
		}
	}

	return nil
}