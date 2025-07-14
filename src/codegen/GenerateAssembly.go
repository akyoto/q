package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

// GenerateAssembly converts the SSA IR to assembler instructions.
func (f *Function) GenerateAssembly(ir ssa.IR, stackFrame bool, hasExternCalls bool) {
	f.createSteps(ir)
	f.Assembler.Append(&asm.Label{Name: f.FullName})

	isInit := f.FullName == "run.init"
	isExit := f.FullName == "os.exit"
	needsFramePointer := !isInit && !isExit

	if f.Preserved.Count() > 0 && !isInit {
		f.Assembler.Append(&asm.PushRegisters{Registers: f.Preserved.Slice()})
	}

	if stackFrame || hasExternCalls {
		f.Assembler.Append(&asm.StackFrameStart{FramePointer: needsFramePointer, ExternCalls: hasExternCalls})
	}

	for _, step := range f.Steps {
		f.exec(step)
	}

	if isExit {
		return
	}

	if stackFrame || hasExternCalls {
		f.Assembler.Append(&asm.StackFrameEnd{FramePointer: needsFramePointer})
	}

	if isInit {
		return
	}

	if f.Preserved.Count() > 0 {
		f.Assembler.Append(&asm.PopRegisters{Registers: f.Preserved.Slice()})
	}

	f.Assembler.Append(&asm.Return{})
}