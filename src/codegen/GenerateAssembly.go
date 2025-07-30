package codegen

import (
	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/x86"
)

// GenerateAssembly converts the SSA IR to assembler instructions.
func (f *Function) GenerateAssembly(ir ssa.IR, build *config.Build, hasStackFrame bool, hasExternCalls bool) {
	f.isInit = f.FullName == "run.init"
	f.isExit = f.FullName == "os.exit"
	f.needsFramePointer = !f.isInit && !f.isExit
	f.hasStackFrame = hasStackFrame
	f.hasExternCalls = hasExternCalls

	switch build.Arch {
	case config.ARM:
		switch build.OS {
		case config.Linux:
			f.CPU = &arm.LinuxCPU
		case config.Mac:
			f.CPU = &arm.MacCPU
		case config.Windows:
			f.CPU = &arm.WindowsCPU
		}
	case config.X86:
		switch build.OS {
		case config.Linux:
			f.CPU = &x86.LinuxCPU
		case config.Mac:
			f.CPU = &x86.MacCPU
		case config.Windows:
			f.CPU = &x86.WindowsCPU
		}
	}

	// Transform SSA graph to a flat slice of steps we have to execute.
	f.createSteps(ir)

	// Execute all steps.
	f.Enter()

	for _, step := range f.Steps {
		f.exec(step)
	}

	if len(f.Steps) == 0 {
		f.Leave()
		return
	}

	_, lastIsReturn := f.Steps[len(f.Steps)-1].Value.(*ssa.Return)

	if !lastIsReturn {
		f.Leave()
	}
}