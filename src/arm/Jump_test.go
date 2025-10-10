package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/go/assert"
)

func TestJump(t *testing.T) {
	usagePatterns := []struct {
		Type   byte
		Offset int
		Code   uint32
	}{
		{0, 0, 0x14000000},
		{0, 1, 0x14000001},
		{0, -1, 0x17FFFFFF},

		{1, 0, 0x54000000},
		{1, 1, 0x54000020},
		{1, -1, 0x54FFFFE0},

		{2, 0, 0x54000001},
		{2, 1, 0x54000021},
		{2, -1, 0x54FFFFE1},

		{3, 0, 0x5400000C},
		{3, 1, 0x5400002C},
		{3, -1, 0x54FFFFEC},

		{4, 0, 0x5400000A},
		{4, 1, 0x5400002A},
		{4, -1, 0x54FFFFEA},

		{5, 0, 0x5400000B},
		{5, 1, 0x5400002B},
		{5, -1, 0x54FFFFEB},

		{6, 0, 0x5400000D},
		{6, 1, 0x5400002D},
		{6, -1, 0x54FFFFED},
	}

	for _, pattern := range usagePatterns {
		t.Logf("b %d", pattern.Offset)

		var (
			code      uint32
			encodable bool
		)

		switch pattern.Type {
		case 0:
			code, encodable = arm.Jump(pattern.Offset)
		case 1:
			code, encodable = arm.JumpIfEqual(pattern.Offset)
		case 2:
			code, encodable = arm.JumpIfNotEqual(pattern.Offset)
		case 3:
			code, encodable = arm.JumpIfGreater(pattern.Offset)
		case 4:
			code, encodable = arm.JumpIfGreaterOrEqual(pattern.Offset)
		case 5:
			code, encodable = arm.JumpIfLess(pattern.Offset)
		case 6:
			code, encodable = arm.JumpIfLessOrEqual(pattern.Offset)
		}

		assert.Equal(t, code, pattern.Code)
		assert.True(t, encodable)
	}
}