package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestDivSignedRegister(t *testing.T) {
	usagePatterns := []struct {
		Register cpu.Register
		Code     []byte
	}{
		{x86.R0, []byte{0x48, 0xF7, 0xF8}},
		{x86.R1, []byte{0x48, 0xF7, 0xF9}},
		{x86.R2, []byte{0x48, 0xF7, 0xFA}},
		{x86.R3, []byte{0x48, 0xF7, 0xFB}},
		{x86.SP, []byte{0x48, 0xF7, 0xFC}},
		{x86.R5, []byte{0x48, 0xF7, 0xFD}},
		{x86.R6, []byte{0x48, 0xF7, 0xFE}},
		{x86.R7, []byte{0x48, 0xF7, 0xFF}},
		{x86.R8, []byte{0x49, 0xF7, 0xF8}},
		{x86.R9, []byte{0x49, 0xF7, 0xF9}},
		{x86.R10, []byte{0x49, 0xF7, 0xFA}},
		{x86.R11, []byte{0x49, 0xF7, 0xFB}},
		{x86.R12, []byte{0x49, 0xF7, 0xFC}},
		{x86.R13, []byte{0x49, 0xF7, 0xFD}},
		{x86.R14, []byte{0x49, 0xF7, 0xFE}},
		{x86.R15, []byte{0x49, 0xF7, 0xFF}},
	}

	for _, pattern := range usagePatterns {
		t.Logf("idiv %s", pattern.Register)
		code := x86.DivSignedRegister(nil, pattern.Register)
		assert.DeepEqual(t, code, pattern.Code)
	}
}

func TestDivUnsignedRegister(t *testing.T) {
	usagePatterns := []struct {
		Register cpu.Register
		Code     []byte
	}{
		{x86.R0, []byte{0x48, 0xF7, 0xF0}},
		{x86.R1, []byte{0x48, 0xF7, 0xF1}},
		{x86.R2, []byte{0x48, 0xF7, 0xF2}},
		{x86.R3, []byte{0x48, 0xF7, 0xF3}},
		{x86.SP, []byte{0x48, 0xF7, 0xF4}},
		{x86.R5, []byte{0x48, 0xF7, 0xF5}},
		{x86.R6, []byte{0x48, 0xF7, 0xF6}},
		{x86.R7, []byte{0x48, 0xF7, 0xF7}},
		{x86.R8, []byte{0x49, 0xF7, 0xF0}},
		{x86.R9, []byte{0x49, 0xF7, 0xF1}},
		{x86.R10, []byte{0x49, 0xF7, 0xF2}},
		{x86.R11, []byte{0x49, 0xF7, 0xF3}},
		{x86.R12, []byte{0x49, 0xF7, 0xF4}},
		{x86.R13, []byte{0x49, 0xF7, 0xF5}},
		{x86.R14, []byte{0x49, 0xF7, 0xF6}},
		{x86.R15, []byte{0x49, 0xF7, 0xF7}},
	}

	for _, pattern := range usagePatterns {
		t.Logf("div %s", pattern.Register)
		code := x86.DivUnsignedRegister(nil, pattern.Register)
		assert.DeepEqual(t, code, pattern.Code)
	}
}