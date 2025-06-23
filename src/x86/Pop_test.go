package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestPopRegister(t *testing.T) {
	usagePatterns := []struct {
		Register cpu.Register
		Code     []byte
	}{
		{x86.R0, []byte{0x58}},
		{x86.R1, []byte{0x59}},
		{x86.R2, []byte{0x5A}},
		{x86.R3, []byte{0x5B}},
		{x86.SP, []byte{0x5C}},
		{x86.R5, []byte{0x5D}},
		{x86.R6, []byte{0x5E}},
		{x86.R7, []byte{0x5F}},
		{x86.R8, []byte{0x41, 0x58}},
		{x86.R9, []byte{0x41, 0x59}},
		{x86.R10, []byte{0x41, 0x5A}},
		{x86.R11, []byte{0x41, 0x5B}},
		{x86.R12, []byte{0x41, 0x5C}},
		{x86.R13, []byte{0x41, 0x5D}},
		{x86.R14, []byte{0x41, 0x5E}},
		{x86.R15, []byte{0x41, 0x5F}},
	}

	for _, pattern := range usagePatterns {
		t.Logf("pop %s", pattern.Register)
		code := x86.PopRegister(nil, pattern.Register)
		assert.DeepEqual(t, code, pattern.Code)
	}
}