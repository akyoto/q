package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestSubRegisterNumber(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Number      int
		Code        uint32
	}{
		{arm.X0, arm.X0, 1, 0xD1000400},
		{arm.X0, arm.X0, 0x1000, 0xD1400400},
		{arm.SP, arm.SP, 16, 0xD10043FF},
	}

	for _, pattern := range usagePatterns {
		t.Logf("sub %s, %s, %d", pattern.Destination, pattern.Source, pattern.Number)
		code, encodable := arm.SubRegisterNumber(pattern.Destination, pattern.Source, pattern.Number)
		assert.True(t, encodable)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestSubRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Operand     cpu.Register
		Code        uint32
	}{
		{arm.X0, arm.X1, arm.X2, 0xCB020020},
	}

	for _, pattern := range usagePatterns {
		t.Logf("sub %s, %s, %s", pattern.Destination, pattern.Source, pattern.Operand)
		code := arm.SubRegisterRegister(pattern.Destination, pattern.Source, pattern.Operand)
		assert.Equal(t, code, pattern.Code)
	}
}