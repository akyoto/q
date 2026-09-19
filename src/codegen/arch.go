package codegen

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

const (
	WindowsTLSOffset = 0x480
	WindowsTLSSize   = 0x200
)

// arch contains the architecture-specific behavior of the codegen phase.
type arch interface {
	// canEncodeNumber returns true if the architecture can encode a number as an immediate for the given instruction.
	canEncodeNumber(instr ssa.Value, number *ssa.Int) bool

	// canStoreNumber returns true if a store instruction can encode the number directly.
	canStoreNumber(typ types.Type, number int) bool

	// casOldValue makes sure the old value of a compare-and-swap is in the architecture's CAS result register.
	casOldValue(f *Function, register cpu.Register) cpu.Register

	// casResultRegister returns the register that holds the old value after a compare-and-swap.
	casResultRegister(register cpu.Register) cpu.Register

	// conditionalSet sets the register to 0 or 1 depending on the condition.
	conditionalSet(f *Function, register cpu.Register, op token.Kind, unsigned bool)

	// conflictsWithStackPointer returns true if the register conflicts with the stack pointer.
	conflictsWithStackPointer(register cpu.Register) bool

	// loadTLS loads the thread-local storage pointer into the destination register.
	loadTLS(f *Function, destination cpu.Register, label string)

	// operandConflict returns true if the register of the given value must differ from the destination register of the instruction.
	operandConflict(instr ssa.Value, value ssa.Value) bool
}

// newArch returns the architecture for the given build.
func newArch(build *config.Build) arch {
	switch build.Arch {
	case config.ARM:
		return archARM{build: build}
	case config.X86:
		return archX86{build: build}
	default:
		panic("unknown architecture")
	}
}