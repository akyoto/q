package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestPushNumber(t *testing.T) {
	usagePatterns := []struct {
		Number int32
		Code   []byte
	}{
		{0, []byte{0x6A, 0x00}},
		{1, []byte{0x6A, 0x01}},
		{-1, []byte{0x6A, 0xFF}},
		{127, []byte{0x6A, 0x7F}},
		{128, []byte{0x68, 0x80, 0x00, 0x00, 0x00}},
		{0xFF, []byte{0x68, 0xFF, 0x00, 0x00, 0x00}},
		{0xFFFF, []byte{0x68, 0xFF, 0xFF, 0x00, 0x00}},
		{0x7FFFFFFF, []byte{0x68, 0xFF, 0xFF, 0xFF, 0x7F}},
	}

	for _, pattern := range usagePatterns {
		t.Logf("push %d", pattern.Number)
		code := x86.PushNumber(nil, pattern.Number)
		assert.DeepEqual(t, code, pattern.Code)
	}
}

func TestPushRegister(t *testing.T) {
	usagePatterns := []struct {
		Register cpu.Register
		Code     []byte
	}{
		{x86.R0, []byte{0x50}},
		{x86.R1, []byte{0x51}},
		{x86.R2, []byte{0x52}},
		{x86.R3, []byte{0x53}},
		{x86.SP, []byte{0x54}},
		{x86.R5, []byte{0x55}},
		{x86.R6, []byte{0x56}},
		{x86.R7, []byte{0x57}},
		{x86.R8, []byte{0x41, 0x50}},
		{x86.R9, []byte{0x41, 0x51}},
		{x86.R10, []byte{0x41, 0x52}},
		{x86.R11, []byte{0x41, 0x53}},
		{x86.R12, []byte{0x41, 0x54}},
		{x86.R13, []byte{0x41, 0x55}},
		{x86.R14, []byte{0x41, 0x56}},
		{x86.R15, []byte{0x41, 0x57}},
	}

	for _, pattern := range usagePatterns {
		t.Logf("push %s", pattern.Register)
		code := x86.PushRegister(nil, pattern.Register)
		assert.DeepEqual(t, code, pattern.Code)
	}
}