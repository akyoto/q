package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestLoadAddress(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Offset      int
		Code        []byte
	}{
		{x86.R0, 0, []byte{0x48, 0x8D, 0x05, 0x00, 0x00, 0x00, 0x00}},
		{x86.R1, 0, []byte{0x48, 0x8D, 0x0D, 0x00, 0x00, 0x00, 0x00}},
		{x86.R2, 0, []byte{0x48, 0x8D, 0x15, 0x00, 0x00, 0x00, 0x00}},
		{x86.R3, 0, []byte{0x48, 0x8D, 0x1D, 0x00, 0x00, 0x00, 0x00}},
		{x86.SP, 0, []byte{0x48, 0x8D, 0x25, 0x00, 0x00, 0x00, 0x00}},
		{x86.R5, 0, []byte{0x48, 0x8D, 0x2D, 0x00, 0x00, 0x00, 0x00}},
		{x86.R6, 0, []byte{0x48, 0x8D, 0x35, 0x00, 0x00, 0x00, 0x00}},
		{x86.R7, 0, []byte{0x48, 0x8D, 0x3D, 0x00, 0x00, 0x00, 0x00}},
		{x86.R8, 0, []byte{0x4C, 0x8D, 0x05, 0x00, 0x00, 0x00, 0x00}},
		{x86.R9, 0, []byte{0x4C, 0x8D, 0x0D, 0x00, 0x00, 0x00, 0x00}},
		{x86.R10, 0, []byte{0x4C, 0x8D, 0x15, 0x00, 0x00, 0x00, 0x00}},
		{x86.R11, 0, []byte{0x4C, 0x8D, 0x1D, 0x00, 0x00, 0x00, 0x00}},
		{x86.R12, 0, []byte{0x4C, 0x8D, 0x25, 0x00, 0x00, 0x00, 0x00}},
		{x86.R13, 0, []byte{0x4C, 0x8D, 0x2D, 0x00, 0x00, 0x00, 0x00}},
		{x86.R14, 0, []byte{0x4C, 0x8D, 0x35, 0x00, 0x00, 0x00, 0x00}},
		{x86.R15, 0, []byte{0x4C, 0x8D, 0x3D, 0x00, 0x00, 0x00, 0x00}},
	}

	for _, pattern := range usagePatterns {
		t.Logf("lea %s, [rip+%d]", pattern.Destination, pattern.Offset)
		code := x86.LoadAddress(nil, pattern.Destination, pattern.Offset)
		assert.DeepEqual(t, code, pattern.Code)
	}
}