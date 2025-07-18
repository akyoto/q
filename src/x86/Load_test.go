package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestLoadRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Base        cpu.Register
		Offset      int8
		Length      byte
		Code        []byte
	}{
		// No offset
		{x86.R0, x86.R15, 0, 8, []byte{0x49, 0x8B, 0x07}},
		{x86.R0, x86.R15, 0, 4, []byte{0x41, 0x8B, 0x07}},
		{x86.R0, x86.R15, 0, 2, []byte{0x66, 0x41, 0x8B, 0x07}},
		{x86.R0, x86.R15, 0, 1, []byte{0x41, 0x8A, 0x07}},
		{x86.R1, x86.R14, 0, 8, []byte{0x49, 0x8B, 0x0E}},
		{x86.R1, x86.R14, 0, 4, []byte{0x41, 0x8B, 0x0E}},
		{x86.R1, x86.R14, 0, 2, []byte{0x66, 0x41, 0x8B, 0x0E}},
		{x86.R1, x86.R14, 0, 1, []byte{0x41, 0x8A, 0x0E}},
		{x86.R2, x86.R13, 0, 8, []byte{0x49, 0x8B, 0x55, 0x00}},
		{x86.R2, x86.R13, 0, 4, []byte{0x41, 0x8B, 0x55, 0x00}},
		{x86.R2, x86.R13, 0, 2, []byte{0x66, 0x41, 0x8B, 0x55, 0x00}},
		{x86.R2, x86.R13, 0, 1, []byte{0x41, 0x8A, 0x55, 0x00}},
		{x86.R3, x86.R12, 0, 8, []byte{0x49, 0x8B, 0x1C, 0x24}},
		{x86.R3, x86.R12, 0, 4, []byte{0x41, 0x8B, 0x1C, 0x24}},
		{x86.R3, x86.R12, 0, 2, []byte{0x66, 0x41, 0x8B, 0x1C, 0x24}},
		{x86.R3, x86.R12, 0, 1, []byte{0x41, 0x8A, 0x1C, 0x24}},
		{x86.SP, x86.R11, 0, 8, []byte{0x49, 0x8B, 0x23}},
		{x86.SP, x86.R11, 0, 4, []byte{0x41, 0x8B, 0x23}},
		{x86.SP, x86.R11, 0, 2, []byte{0x66, 0x41, 0x8B, 0x23}},
		{x86.SP, x86.R11, 0, 1, []byte{0x41, 0x8A, 0x23}},
		{x86.R5, x86.R10, 0, 8, []byte{0x49, 0x8B, 0x2A}},
		{x86.R5, x86.R10, 0, 4, []byte{0x41, 0x8B, 0x2A}},
		{x86.R5, x86.R10, 0, 2, []byte{0x66, 0x41, 0x8B, 0x2A}},
		{x86.R5, x86.R10, 0, 1, []byte{0x41, 0x8A, 0x2A}},
		{x86.R6, x86.R9, 0, 8, []byte{0x49, 0x8B, 0x31}},
		{x86.R6, x86.R9, 0, 4, []byte{0x41, 0x8B, 0x31}},
		{x86.R6, x86.R9, 0, 2, []byte{0x66, 0x41, 0x8B, 0x31}},
		{x86.R6, x86.R9, 0, 1, []byte{0x41, 0x8A, 0x31}},
		{x86.R7, x86.R8, 0, 8, []byte{0x49, 0x8B, 0x38}},
		{x86.R7, x86.R8, 0, 4, []byte{0x41, 0x8B, 0x38}},
		{x86.R7, x86.R8, 0, 2, []byte{0x66, 0x41, 0x8B, 0x38}},
		{x86.R7, x86.R8, 0, 1, []byte{0x41, 0x8A, 0x38}},
		{x86.R8, x86.R7, 0, 8, []byte{0x4C, 0x8B, 0x07}},
		{x86.R8, x86.R7, 0, 4, []byte{0x44, 0x8B, 0x07}},
		{x86.R8, x86.R7, 0, 2, []byte{0x66, 0x44, 0x8B, 0x07}},
		{x86.R8, x86.R7, 0, 1, []byte{0x44, 0x8A, 0x07}},
		{x86.R9, x86.R6, 0, 8, []byte{0x4C, 0x8B, 0x0E}},
		{x86.R9, x86.R6, 0, 4, []byte{0x44, 0x8B, 0x0E}},
		{x86.R9, x86.R6, 0, 2, []byte{0x66, 0x44, 0x8B, 0x0E}},
		{x86.R9, x86.R6, 0, 1, []byte{0x44, 0x8A, 0x0E}},
		{x86.R10, x86.R5, 0, 8, []byte{0x4C, 0x8B, 0x55, 0x00}},
		{x86.R10, x86.R5, 0, 4, []byte{0x44, 0x8B, 0x55, 0x00}},
		{x86.R10, x86.R5, 0, 2, []byte{0x66, 0x44, 0x8B, 0x55, 0x00}},
		{x86.R10, x86.R5, 0, 1, []byte{0x44, 0x8A, 0x55, 0x00}},
		{x86.R11, x86.SP, 0, 8, []byte{0x4C, 0x8B, 0x1C, 0x24}},
		{x86.R11, x86.SP, 0, 4, []byte{0x44, 0x8B, 0x1C, 0x24}},
		{x86.R11, x86.SP, 0, 2, []byte{0x66, 0x44, 0x8B, 0x1C, 0x24}},
		{x86.R11, x86.SP, 0, 1, []byte{0x44, 0x8A, 0x1C, 0x24}},
		{x86.R12, x86.R3, 0, 8, []byte{0x4C, 0x8B, 0x23}},
		{x86.R12, x86.R3, 0, 4, []byte{0x44, 0x8B, 0x23}},
		{x86.R12, x86.R3, 0, 2, []byte{0x66, 0x44, 0x8B, 0x23}},
		{x86.R12, x86.R3, 0, 1, []byte{0x44, 0x8A, 0x23}},
		{x86.R13, x86.R2, 0, 8, []byte{0x4C, 0x8B, 0x2A}},
		{x86.R13, x86.R2, 0, 4, []byte{0x44, 0x8B, 0x2A}},
		{x86.R13, x86.R2, 0, 2, []byte{0x66, 0x44, 0x8B, 0x2A}},
		{x86.R13, x86.R2, 0, 1, []byte{0x44, 0x8A, 0x2A}},
		{x86.R14, x86.R1, 0, 8, []byte{0x4C, 0x8B, 0x31}},
		{x86.R14, x86.R1, 0, 4, []byte{0x44, 0x8B, 0x31}},
		{x86.R14, x86.R1, 0, 2, []byte{0x66, 0x44, 0x8B, 0x31}},
		{x86.R14, x86.R1, 0, 1, []byte{0x44, 0x8A, 0x31}},
		{x86.R15, x86.R0, 0, 8, []byte{0x4C, 0x8B, 0x38}},
		{x86.R15, x86.R0, 0, 4, []byte{0x44, 0x8B, 0x38}},
		{x86.R15, x86.R0, 0, 2, []byte{0x66, 0x44, 0x8B, 0x38}},
		{x86.R15, x86.R0, 0, 1, []byte{0x44, 0x8A, 0x38}},

		// Offset of 1
		{x86.R0, x86.R15, 1, 8, []byte{0x49, 0x8B, 0x47, 0x01}},
		{x86.R0, x86.R15, 1, 4, []byte{0x41, 0x8B, 0x47, 0x01}},
		{x86.R0, x86.R15, 1, 2, []byte{0x66, 0x41, 0x8B, 0x47, 0x01}},
		{x86.R0, x86.R15, 1, 1, []byte{0x41, 0x8A, 0x47, 0x01}},
		{x86.R1, x86.R14, 1, 8, []byte{0x49, 0x8B, 0x4E, 0x01}},
		{x86.R1, x86.R14, 1, 4, []byte{0x41, 0x8B, 0x4E, 0x01}},
		{x86.R1, x86.R14, 1, 2, []byte{0x66, 0x41, 0x8B, 0x4E, 0x01}},
		{x86.R1, x86.R14, 1, 1, []byte{0x41, 0x8A, 0x4E, 0x01}},
		{x86.R2, x86.R13, 1, 8, []byte{0x49, 0x8B, 0x55, 0x01}},
		{x86.R2, x86.R13, 1, 4, []byte{0x41, 0x8B, 0x55, 0x01}},
		{x86.R2, x86.R13, 1, 2, []byte{0x66, 0x41, 0x8B, 0x55, 0x01}},
		{x86.R2, x86.R13, 1, 1, []byte{0x41, 0x8A, 0x55, 0x01}},
		{x86.R3, x86.R12, 1, 8, []byte{0x49, 0x8B, 0x5C, 0x24, 0x01}},
		{x86.R3, x86.R12, 1, 4, []byte{0x41, 0x8B, 0x5C, 0x24, 0x01}},
		{x86.R3, x86.R12, 1, 2, []byte{0x66, 0x41, 0x8B, 0x5C, 0x24, 0x01}},
		{x86.R3, x86.R12, 1, 1, []byte{0x41, 0x8A, 0x5C, 0x24, 0x01}},
		{x86.SP, x86.R11, 1, 8, []byte{0x49, 0x8B, 0x63, 0x01}},
		{x86.SP, x86.R11, 1, 4, []byte{0x41, 0x8B, 0x63, 0x01}},
		{x86.SP, x86.R11, 1, 2, []byte{0x66, 0x41, 0x8B, 0x63, 0x01}},
		{x86.SP, x86.R11, 1, 1, []byte{0x41, 0x8A, 0x63, 0x01}},
		{x86.R5, x86.R10, 1, 8, []byte{0x49, 0x8B, 0x6A, 0x01}},
		{x86.R5, x86.R10, 1, 4, []byte{0x41, 0x8B, 0x6A, 0x01}},
		{x86.R5, x86.R10, 1, 2, []byte{0x66, 0x41, 0x8B, 0x6A, 0x01}},
		{x86.R5, x86.R10, 1, 1, []byte{0x41, 0x8A, 0x6A, 0x01}},
		{x86.R6, x86.R9, 1, 8, []byte{0x49, 0x8B, 0x71, 0x01}},
		{x86.R6, x86.R9, 1, 4, []byte{0x41, 0x8B, 0x71, 0x01}},
		{x86.R6, x86.R9, 1, 2, []byte{0x66, 0x41, 0x8B, 0x71, 0x01}},
		{x86.R6, x86.R9, 1, 1, []byte{0x41, 0x8A, 0x71, 0x01}},
		{x86.R7, x86.R8, 1, 8, []byte{0x49, 0x8B, 0x78, 0x01}},
		{x86.R7, x86.R8, 1, 4, []byte{0x41, 0x8B, 0x78, 0x01}},
		{x86.R7, x86.R8, 1, 2, []byte{0x66, 0x41, 0x8B, 0x78, 0x01}},
		{x86.R7, x86.R8, 1, 1, []byte{0x41, 0x8A, 0x78, 0x01}},
		{x86.R8, x86.R7, 1, 8, []byte{0x4C, 0x8B, 0x47, 0x01}},
		{x86.R8, x86.R7, 1, 4, []byte{0x44, 0x8B, 0x47, 0x01}},
		{x86.R8, x86.R7, 1, 2, []byte{0x66, 0x44, 0x8B, 0x47, 0x01}},
		{x86.R8, x86.R7, 1, 1, []byte{0x44, 0x8A, 0x47, 0x01}},
		{x86.R9, x86.R6, 1, 8, []byte{0x4C, 0x8B, 0x4E, 0x01}},
		{x86.R9, x86.R6, 1, 4, []byte{0x44, 0x8B, 0x4E, 0x01}},
		{x86.R9, x86.R6, 1, 2, []byte{0x66, 0x44, 0x8B, 0x4E, 0x01}},
		{x86.R9, x86.R6, 1, 1, []byte{0x44, 0x8A, 0x4E, 0x01}},
		{x86.R10, x86.R5, 1, 8, []byte{0x4C, 0x8B, 0x55, 0x01}},
		{x86.R10, x86.R5, 1, 4, []byte{0x44, 0x8B, 0x55, 0x01}},
		{x86.R10, x86.R5, 1, 2, []byte{0x66, 0x44, 0x8B, 0x55, 0x01}},
		{x86.R10, x86.R5, 1, 1, []byte{0x44, 0x8A, 0x55, 0x01}},
		{x86.R11, x86.SP, 1, 8, []byte{0x4C, 0x8B, 0x5C, 0x24, 0x01}},
		{x86.R11, x86.SP, 1, 4, []byte{0x44, 0x8B, 0x5C, 0x24, 0x01}},
		{x86.R11, x86.SP, 1, 2, []byte{0x66, 0x44, 0x8B, 0x5C, 0x24, 0x01}},
		{x86.R11, x86.SP, 1, 1, []byte{0x44, 0x8A, 0x5C, 0x24, 0x01}},
		{x86.R12, x86.R3, 1, 8, []byte{0x4C, 0x8B, 0x63, 0x01}},
		{x86.R12, x86.R3, 1, 4, []byte{0x44, 0x8B, 0x63, 0x01}},
		{x86.R12, x86.R3, 1, 2, []byte{0x66, 0x44, 0x8B, 0x63, 0x01}},
		{x86.R12, x86.R3, 1, 1, []byte{0x44, 0x8A, 0x63, 0x01}},
		{x86.R13, x86.R2, 1, 8, []byte{0x4C, 0x8B, 0x6A, 0x01}},
		{x86.R13, x86.R2, 1, 4, []byte{0x44, 0x8B, 0x6A, 0x01}},
		{x86.R13, x86.R2, 1, 2, []byte{0x66, 0x44, 0x8B, 0x6A, 0x01}},
		{x86.R13, x86.R2, 1, 1, []byte{0x44, 0x8A, 0x6A, 0x01}},
		{x86.R14, x86.R1, 1, 8, []byte{0x4C, 0x8B, 0x71, 0x01}},
		{x86.R14, x86.R1, 1, 4, []byte{0x44, 0x8B, 0x71, 0x01}},
		{x86.R14, x86.R1, 1, 2, []byte{0x66, 0x44, 0x8B, 0x71, 0x01}},
		{x86.R14, x86.R1, 1, 1, []byte{0x44, 0x8A, 0x71, 0x01}},
		{x86.R15, x86.R0, 1, 8, []byte{0x4C, 0x8B, 0x78, 0x01}},
		{x86.R15, x86.R0, 1, 4, []byte{0x44, 0x8B, 0x78, 0x01}},
		{x86.R15, x86.R0, 1, 2, []byte{0x66, 0x44, 0x8B, 0x78, 0x01}},
		{x86.R15, x86.R0, 1, 1, []byte{0x44, 0x8A, 0x78, 0x01}},
	}

	for _, pattern := range usagePatterns {
		t.Logf("load %dB %s, [%s+%d]", pattern.Length, pattern.Destination, pattern.Base, pattern.Offset)
		code := x86.LoadRegister(nil, pattern.Destination, pattern.Base, pattern.Offset, pattern.Length)
		assert.DeepEqual(t, code, pattern.Code)
	}
}