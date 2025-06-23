package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestJump(t *testing.T) {
	usagePatterns := []struct {
		Offset int8
		Code   []byte
	}{
		{0, []byte{0xEB, 0x00}},
		{1, []byte{0xEB, 0x01}},
		{2, []byte{0xEB, 0x02}},
		{3, []byte{0xEB, 0x03}},
		{127, []byte{0xEB, 0x7F}},
		{-1, []byte{0xEB, 0xFF}},
		{-2, []byte{0xEB, 0xFE}},
		{-3, []byte{0xEB, 0xFD}},
		{-128, []byte{0xEB, 0x80}},
	}

	for _, pattern := range usagePatterns {
		t.Logf("jmp %x", pattern.Offset)
		code := x86.Jump8(nil, pattern.Offset)
		assert.DeepEqual(t, code, pattern.Code)
	}
}

func TestConditionalJump(t *testing.T) {
	assert.DeepEqual(t, x86.Jump8IfEqual(nil, 1), []byte{0x74, 0x01})
	assert.DeepEqual(t, x86.Jump8IfNotEqual(nil, 1), []byte{0x75, 0x01})
	assert.DeepEqual(t, x86.Jump8IfLess(nil, 1), []byte{0x7C, 0x01})
	assert.DeepEqual(t, x86.Jump8IfGreaterOrEqual(nil, 1), []byte{0x7D, 0x01})
	assert.DeepEqual(t, x86.Jump8IfLessOrEqual(nil, 1), []byte{0x7E, 0x01})
	assert.DeepEqual(t, x86.Jump8IfGreater(nil, 1), []byte{0x7F, 0x01})
}