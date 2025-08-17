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
		General: []cpu.Register{
			R1, R2, R6, R7, R8, R9, R10, R11, // Clobbered
			R3, R12, R13, R14, R15, // Preserved
		},
		Call: cpu.ABI{
			In:        []cpu.Register{R0, R7, R6, R2, R10, R8, R9},
			Out:       []cpu.Register{R0, R7, R6, R2, R10, R8, R9},
			Clobbered: []cpu.Register{R0, R1, R2, R6, R7, R8, R9, R10, R11},
			Preserved: []cpu.Register{R3, SP, R5, R12, R13, R14, R15},
		},
		ExternCall: cpu.ABI{
			In:        []cpu.Register{R7, R6, R2, R1, R8, R9},
			Out:       []cpu.Register{R0, R2},
			Clobbered: []cpu.Register{R0, R1, R2, R6, R7, R8, R9, R10, R11},
			Preserved: []cpu.Register{R3, SP, R5, R12, R13, R14, R15},
		},
		Syscall: cpu.ABI{
			In:        []cpu.Register{R0, R7, R6, R2, R10, R8, R9},
			Out:       []cpu.Register{R0},
			Clobbered: []cpu.Register{R0, R1, R11},
			Preserved: []cpu.Register{R2, R3, SP, R5, R6, R7, R8, R9, R10, R12, R13, R14, R15},
		},
		DivisionClobbered: []cpu.Register{R0, R2},
		DivisorRestricted: []cpu.Register{R2},
		ShiftClobbered:    []cpu.Register{R1},
		ShiftRestricted:   []cpu.Register{R1},
	}

	MacCPU = cpu.CPU{
		General:    LinuxCPU.General,
		Call:       LinuxCPU.Call,
		ExternCall: LinuxCPU.ExternCall,
		Syscall: cpu.ABI{
			In:        LinuxCPU.Syscall.In,
			Out:       LinuxCPU.Syscall.Out,
			Clobbered: []cpu.Register{R0, R1, R2, R11},
			Preserved: []cpu.Register{R3, SP, R5, R6, R7, R8, R9, R10, R12, R13, R14, R15},
		},
		DivisionClobbered: LinuxCPU.DivisionClobbered,
		DivisorRestricted: LinuxCPU.DivisorRestricted,
		ShiftClobbered:    LinuxCPU.ShiftClobbered,
		ShiftRestricted:   LinuxCPU.ShiftRestricted,
	}

	WindowsCPU = cpu.CPU{
		General: []cpu.Register{
			R1, R2, R8, R9, R10, R11, // Clobbered
			R3, R6, R7, R12, R13, R14, R15, // Preserved
		},
		Call: LinuxCPU.Call,
		ExternCall: cpu.ABI{
			In:        []cpu.Register{R1, R2, R8, R9},
			Out:       []cpu.Register{R0},
			Clobbered: []cpu.Register{R0, R1, R2, R8, R9, R10, R11},
			Preserved: []cpu.Register{R3, SP, R5, R6, R7, R12, R13, R14, R15},
		},
		DivisionClobbered: LinuxCPU.DivisionClobbered,
		DivisorRestricted: LinuxCPU.DivisorRestricted,
		ShiftClobbered:    LinuxCPU.ShiftClobbered,
		ShiftRestricted:   LinuxCPU.ShiftRestricted,
	}
)