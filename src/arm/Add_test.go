package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestAddRegisterNumber(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Number      int
		Code        uint32
	}{
		{arm.X0, arm.X0, 1, 0x91000400},
		{arm.X0, arm.X0, 0x1000, 0x91400400},
	}

	for _, pattern := range usagePatterns {
		t.Logf("add %s, %s, %d", pattern.Destination, pattern.Source, pattern.Number)
		code, encodable := arm.AddRegisterNumber(pattern.Destination, pattern.Source, pattern.Number)
		assert.True(t, encodable)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestAddRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Operand     cpu.Register
		Code        uint32
	}{
		{arm.X0, arm.X1, arm.X2, 0x8B020020},
	}

	for _, pattern := range usagePatterns {
		t.Logf("add %s, %s, %s", pattern.Destination, pattern.Source, pattern.Operand)
		code := arm.AddRegisterRegister(pattern.Destination, pattern.Source, pattern.Operand)
		assert.Equal(t, code, pattern.Code)
	}
}