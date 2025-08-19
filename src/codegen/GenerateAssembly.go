package codegen

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/ssa"
)

// GenerateAssembly converts the SSA IR to assembler instructions.
func (f *Function) GenerateAssembly(ir ssa.IR, build *config.Build, hasStackFrame bool, hasExternCalls bool) {
	f.isInit = f.FullName == "run.init"
	f.isExit = f.FullName == "os.exit"
	f.needsFramePointer = !f.isInit && !f.isExit
	f.hasStackFrame = hasStackFrame
	f.hasExternCalls = hasExternCalls
	f.build = build
	f.CPU = selectCPU(build)

	// Transform SSA graph to a flat slice of steps we have to execute.
	f.createSteps(ir)

	// Execute all steps.
	f.enter()

	for _, step := range f.Steps {
		f.exec(step)
	}

	if len(f.Steps) == 0 {
		f.leave()
		return
	}

	_, lastIsReturn := f.Steps[len(f.Steps)-1].Value.(*ssa.Return)

	if !lastIsReturn {
		f.leave()
	}
}