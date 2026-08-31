package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestSignExtend(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Length      byte
		Code        uint32
	}{
		{arm.X0, arm.X0, 1, 0x93401C00},
		{arm.X3, arm.X2, 1, 0x93401C43},
		{arm.X15, arm.X8, 1, 0x93401D0F},
		{arm.X0, arm.X0, 2, 0x93403C00},
		{arm.X3, arm.X2, 2, 0x93403C43},
		{arm.ZR, arm.ZR, 2, 0x93403FFF},
		{arm.X0, arm.X0, 4, 0x93407C00},
		{arm.X15, arm.X8, 4, 0x93407D0F},
		{arm.ZR, arm.ZR, 4, 0x93407FFF},
	}

	for _, pattern := range usagePatterns {
		t.Logf("sxt %db %s, %s", pattern.Length, pattern.Destination, pattern.Source)
		code := arm.SignExtend(pattern.Destination, pattern.Source, pattern.Length)
		assert.Equal(t, code, pattern.Code)
	}
}