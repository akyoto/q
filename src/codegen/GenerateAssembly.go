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

	// Transform SSA graph to a flat slice of steps we can execute one by one.
	f.createSteps(ir)

	// Execute all steps to produce assembly code.
	f.generate()
}