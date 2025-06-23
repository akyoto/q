package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestNegateRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Code        uint32
	}{
		{arm.X0, arm.X0, 0xCB0003E0},
		{arm.X1, arm.X1, 0xCB0103E1},
	}

	for _, pattern := range usagePatterns {
		t.Logf("neg %s, %s", pattern.Destination, pattern.Source)
		code := arm.NegateRegister(pattern.Destination, pattern.Source)
		assert.Equal(t, code, pattern.Code)
	}
}