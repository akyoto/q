package x86

import "git.urbach.dev/cli/q/src/cpu"

const (
	R0 cpu.Register = iota // RAX
	R1                     // RCX
	R2                     // RDX
	R3                     // RBX
	SP                     // Stack pointer
	R5                     // RBP
	R6                     // RSI
	R7                     // RDI
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
)

var CPU = cpu.CPU{
	Call:               []cpu.Register{R0, R7, R6, R2, R10, R8, R9},
	Syscall:            []cpu.Register{R0, R7, R6, R2, R10, R8, R9},
	ExternCall:         []cpu.Register{R1, R2, R8, R9},
	ExternCallVolatile: []cpu.Register{R0, R1, R2, R8, R9, R10, R11},
	Return:             []cpu.Register{R0, R7, R6},
}