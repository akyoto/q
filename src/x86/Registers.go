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

var (
	LinuxCPU = cpu.CPU{
		Call: cpu.ABI{
			In:       []cpu.Register{R0, R7, R6, R2, R10, R8, R9},
			Out:      []cpu.Register{R0, R7, R6},
			Volatile: []cpu.Register{R0, R1, R2, R6, R7, R8, R9, R10, R11},
		},
		ExternCall: cpu.ABI{
			In:       []cpu.Register{R7, R6, R2, R1, R8, R9},
			Out:      []cpu.Register{R0, R2},
			Volatile: []cpu.Register{R0, R1, R2, R6, R7, R8, R9, R10, R11},
		},
		Syscall: cpu.ABI{
			In:       []cpu.Register{R0, R7, R6, R2, R10, R8, R9},
			Out:      []cpu.Register{R0},
			Volatile: []cpu.Register{R0, R7, R6},
		},
	}

	MacCPU = LinuxCPU

	WindowsCPU = cpu.CPU{
		Call: LinuxCPU.Call,
		ExternCall: cpu.ABI{
			In:       []cpu.Register{R1, R2, R8, R9},
			Out:      []cpu.Register{R0},
			Volatile: []cpu.Register{R0, R1, R2, R8, R9, R10, R11},
		},
	}
)