package codegen

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/ssa"
)

// CompileToAssembly converts the SSA IR to assembler instructions.
func (f *Function) CompileToAssembly(ir ssa.IR, build *config.Build, hasStackFrame bool, hasExternCalls bool) {
	f.isInit = f.FullName == "run.init"
	f.isExit = f.FullName == "run.exit"
	f.needsFramePointer = !f.isInit && !f.isExit
	f.hasStackFrame = hasStackFrame
	f.hasExternCalls = hasExternCalls
	f.build = build

	// Transform SSA graph to a flat slice of steps we can execute one by one.
	f.createSteps(ir)

	// Create the stack frame and preserve registers if needed.
	f.enter()

	// Execute all steps to produce assembly code.
	for _, step := range f.Steps {
		f.execute(step)
	}
}