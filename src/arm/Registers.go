package arm

import "git.urbach.dev/cli/q/src/cpu"

const (
	X0 cpu.Register = iota // Function arguments and return values [0-7]
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8 // Indirect result location register (used to pass a pointer to a structure return value)
	X9 // Temporary registers (caller-saved, used for general computation) [9-15]
	X10
	X11
	X12
	X13
	X14
	X15
	X16 // Intra-procedure call scratch registers [16-17]
	X17
	X18 // Platform register (reserved by the platform ABI for thread-local storage)
	X19 // Callee-saved registers (must be preserved across function calls) [19-28]
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27
	X28
	FP // Frame pointer
	LR // Link register
	SP // Stack pointer
)

const (
	ZR = SP // Zero register uses the same numerical value as SP
)

var CPU = cpu.CPU{
	Call:       []cpu.Register{X0, X1, X2, X3, X4, X5, X6},
	Syscall:    []cpu.Register{X8, X0, X1, X2, X3, X4, X5},
	ExternCall: []cpu.Register{X0, X1, X2, X3, X4, X5, X6, X7},
}
