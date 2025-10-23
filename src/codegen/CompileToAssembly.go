package codegen

import (
	"slices"

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
	f.IR = createSteps(ir)
	f.reorderPhis()

	for _, step := range slices.Backward(f.Steps) {
		f.hintABI(step)
		f.createLiveRanges(step)
	}

	for _, step := range slices.Backward(f.Steps) {
		if step.Register == -1 && f.needsRegister(step) {
			f.assignFreeRegister(step)
		}

		f.hintDestination(step)
	}

	f.reorderParameters()
	f.fixRegisterConflicts()
	f.addPreservedRegisters()
	f.enter()

	for _, step := range f.Steps {
		f.execute(step)
	}
}