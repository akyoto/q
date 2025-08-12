package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestLoadAddress(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Number      int
		Code        uint32
	}{
		{arm.X0, 56, 0x100001C0},
		{arm.X1, 80, 0x10000281},
		{arm.X16, 0x3000, 0x10018010},
	}

	for _, pattern := range usagePatterns {
		t.Logf("adr %s, %d", pattern.Destination, pattern.Number)
		code := arm.LoadAddress(pattern.Destination, pattern.Number)
		assert.Equal(t, code, pattern.Code)
	}
}