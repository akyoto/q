package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestAndRegisterNumber(t *testing.T) {
	usagePatterns := []struct {
		Register cpu.Register
		Number   int
		Code     []byte
	}{
		{x86.R0, 1, []byte{0x48, 0x83, 0xE0, 0x01}},
		{x86.R1, 1, []byte{0x48, 0x83, 0xE1, 0x01}},
		{x86.R2, 1, []byte{0x48, 0x83, 0xE2, 0x01}},
		{x86.R3, 1, []byte{0x48, 0x83, 0xE3, 0x01}},
		{x86.SP, 1, []byte{0x48, 0x83, 0xE4, 0x01}},
		{x86.R5, 1, []byte{0x48, 0x83, 0xE5, 0x01}},
		{x86.R6, 1, []byte{0x48, 0x83, 0xE6, 0x01}},
		{x86.R7, 1, []byte{0x48, 0x83, 0xE7, 0x01}},
		{x86.R8, 1, []byte{0x49, 0x83, 0xE0, 0x01}},
		{x86.R9, 1, []byte{0x49, 0x83, 0xE1, 0x01}},
		{x86.R10, 1, []byte{0x49, 0x83, 0xE2, 0x01}},
		{x86.R11, 1, []byte{0x49, 0x83, 0xE3, 0x01}},
		{x86.R12, 1, []byte{0x49, 0x83, 0xE4, 0x01}},
		{x86.R13, 1, []byte{0x49, 0x83, 0xE5, 0x01}},
		{x86.R14, 1, []byte{0x49, 0x83, 0xE6, 0x01}},
		{x86.R15, 1, []byte{0x49, 0x83, 0xE7, 0x01}},

		{x86.R0, 0x7FFFFFFF, []byte{0x48, 0x81, 0xE0, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R1, 0x7FFFFFFF, []byte{0x48, 0x81, 0xE1, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R2, 0x7FFFFFFF, []byte{0x48, 0x81, 0xE2, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R3, 0x7FFFFFFF, []byte{0x48, 0x81, 0xE3, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.SP, 0x7FFFFFFF, []byte{0x48, 0x81, 0xE4, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R5, 0x7FFFFFFF, []byte{0x48, 0x81, 0xE5, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R6, 0x7FFFFFFF, []byte{0x48, 0x81, 0xE6, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R7, 0x7FFFFFFF, []byte{0x48, 0x81, 0xE7, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R8, 0x7FFFFFFF, []byte{0x49, 0x81, 0xE0, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R9, 0x7FFFFFFF, []byte{0x49, 0x81, 0xE1, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R10, 0x7FFFFFFF, []byte{0x49, 0x81, 0xE2, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R11, 0x7FFFFFFF, []byte{0x49, 0x81, 0xE3, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R12, 0x7FFFFFFF, []byte{0x49, 0x81, 0xE4, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R13, 0x7FFFFFFF, []byte{0x49, 0x81, 0xE5, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R14, 0x7FFFFFFF, []byte{0x49, 0x81, 0xE6, 0xFF, 0xFF, 0xFF, 0x7F}},
		{x86.R15, 0x7FFFFFFF, []byte{0x49, 0x81, 0xE7, 0xFF, 0xFF, 0xFF, 0x7F}},
	}

	for _, pattern := range usagePatterns {
		t.Logf("and %s, %x", pattern.Register, pattern.Number)
		code := x86.AndRegisterNumber(nil, pattern.Register, pattern.Number)
		assert.DeepEqual(t, code, pattern.Code)
	}
}

func TestAndRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Left  cpu.Register
		Right cpu.Register
		Code  []byte
	}{
		{x86.R0, x86.R15, []byte{0x4C, 0x21, 0xF8}},
		{x86.R1, x86.R14, []byte{0x4C, 0x21, 0xF1}},
		{x86.R2, x86.R13, []byte{0x4C, 0x21, 0xEA}},
		{x86.R3, x86.R12, []byte{0x4C, 0x21, 0xE3}},
		{x86.SP, x86.R11, []byte{0x4C, 0x21, 0xDC}},
		{x86.R5, x86.R10, []byte{0x4C, 0x21, 0xD5}},
		{x86.R6, x86.R9, []byte{0x4C, 0x21, 0xCE}},
		{x86.R7, x86.R8, []byte{0x4C, 0x21, 0xC7}},
		{x86.R8, x86.R7, []byte{0x49, 0x21, 0xF8}},
		{x86.R9, x86.R6, []byte{0x49, 0x21, 0xF1}},
		{x86.R10, x86.R5, []byte{0x49, 0x21, 0xEA}},
		{x86.R11, x86.SP, []byte{0x49, 0x21, 0xE3}},
		{x86.R12, x86.R3, []byte{0x49, 0x21, 0xDC}},
		{x86.R13, x86.R2, []byte{0x49, 0x21, 0xD5}},
		{x86.R14, x86.R1, []byte{0x49, 0x21, 0xCE}},
		{x86.R15, x86.R0, []byte{0x49, 0x21, 0xC7}},
	}

	for _, pattern := range usagePatterns {
		t.Logf("and %s, %s", pattern.Left, pattern.Right)
		code := x86.AndRegisterRegister(nil, pattern.Left, pattern.Right)
		assert.DeepEqual(t, code, pattern.Code)
	}
}