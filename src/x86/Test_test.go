package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestTestRegister(t *testing.T) {
	usagePatterns := []struct {
		Left  cpu.Register
		Right cpu.Register
		Code  []byte
	}{
		{x86.R0, x86.R0, []byte{0x48, 0x85, 0xC0}},
		{x86.R1, x86.R1, []byte{0x48, 0x85, 0xC9}},
		{x86.R2, x86.R2, []byte{0x48, 0x85, 0xD2}},
		{x86.R3, x86.R3, []byte{0x48, 0x85, 0xDB}},
		{x86.SP, x86.SP, []byte{0x48, 0x85, 0xE4}},
		{x86.R5, x86.R5, []byte{0x48, 0x85, 0xED}},
		{x86.R6, x86.R6, []byte{0x48, 0x85, 0xF6}},
		{x86.R7, x86.R7, []byte{0x48, 0x85, 0xFF}},
		{x86.R8, x86.R8, []byte{0x4D, 0x85, 0xC0}},
		{x86.R9, x86.R9, []byte{0x4D, 0x85, 0xC9}},
		{x86.R10, x86.R10, []byte{0x4D, 0x85, 0xD2}},
		{x86.R11, x86.R11, []byte{0x4D, 0x85, 0xDB}},
		{x86.R12, x86.R12, []byte{0x4D, 0x85, 0xE4}},
		{x86.R13, x86.R13, []byte{0x4D, 0x85, 0xED}},
		{x86.R14, x86.R14, []byte{0x4D, 0x85, 0xF6}},
		{x86.R15, x86.R15, []byte{0x4D, 0x85, 0xFF}},

		{x86.R0, x86.R15, []byte{0x4C, 0x85, 0xF8}},
		{x86.R1, x86.R14, []byte{0x4C, 0x85, 0xF1}},
		{x86.R2, x86.R13, []byte{0x4C, 0x85, 0xEA}},
		{x86.R3, x86.R12, []byte{0x4C, 0x85, 0xE3}},
		{x86.SP, x86.R11, []byte{0x4C, 0x85, 0xDC}},
		{x86.R5, x86.R10, []byte{0x4C, 0x85, 0xD5}},
		{x86.R6, x86.R9, []byte{0x4C, 0x85, 0xCE}},
		{x86.R7, x86.R8, []byte{0x4C, 0x85, 0xC7}},
		{x86.R8, x86.R7, []byte{0x49, 0x85, 0xF8}},
		{x86.R9, x86.R6, []byte{0x49, 0x85, 0xF1}},
		{x86.R10, x86.R5, []byte{0x49, 0x85, 0xEA}},
		{x86.R11, x86.SP, []byte{0x49, 0x85, 0xE3}},
		{x86.R12, x86.R3, []byte{0x49, 0x85, 0xDC}},
		{x86.R13, x86.R2, []byte{0x49, 0x85, 0xD5}},
		{x86.R14, x86.R1, []byte{0x49, 0x85, 0xCE}},
		{x86.R15, x86.R0, []byte{0x49, 0x85, 0xC7}},
	}

	for _, pattern := range usagePatterns {
		t.Logf("test %s, %s", pattern.Left, pattern.Right)
		code := x86.TestRegisterRegister(nil, pattern.Left, pattern.Right)
		assert.DeepEqual(t, code, pattern.Code)
	}
}