package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestSet(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Condition   byte
		Code        uint32
	}{
		{arm.X0, 1, 0x9A9F17E0},
		{arm.X1, 1, 0x9A9F17E1},
		{arm.X0, 2, 0x9A9F07E0},
		{arm.X1, 2, 0x9A9F07E1},
		{arm.X0, 3, 0x9A9FD7E0},
		{arm.X1, 3, 0x9A9FD7E1},
		{arm.X0, 4, 0x9A9FB7E0},
		{arm.X1, 4, 0x9A9FB7E1},
		{arm.X0, 5, 0x9A9FA7E0},
		{arm.X1, 5, 0x9A9FA7E1},
		{arm.X0, 6, 0x9A9FC7E0},
		{arm.X1, 6, 0x9A9FC7E1},
	}

	for _, pattern := range usagePatterns {
		var code uint32

		switch pattern.Condition {
		case 1:
			t.Logf("cset %s, eq", pattern.Destination)
			code = arm.SetIfEqual(pattern.Destination)
		case 2:
			t.Logf("cset %s, ne", pattern.Destination)
			code = arm.SetIfNotEqual(pattern.Destination)
		case 3:
			t.Logf("cset %s, gt", pattern.Destination)
			code = arm.SetIfGreater(pattern.Destination)
		case 4:
			t.Logf("cset %s, ge", pattern.Destination)
			code = arm.SetIfGreaterOrEqual(pattern.Destination)
		case 5:
			t.Logf("cset %s, lt", pattern.Destination)
			code = arm.SetIfLess(pattern.Destination)
		case 6:
			t.Logf("cset %s, le", pattern.Destination)
			code = arm.SetIfLessOrEqual(pattern.Destination)
		}

		assert.Equal(t, code, pattern.Code)
	}
}