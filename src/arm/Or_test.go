package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestOrRegisterNumber(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Number      int
		Code        uint32
	}{
		{arm.X0, arm.X1, 1, 0xB2400020},
		{arm.X0, arm.X1, 2, 0xB27F0020},
		{arm.X0, arm.X1, 3, 0xB2400420},
		{arm.X0, arm.X1, 7, 0xB2400820},
		{arm.X0, arm.X1, 16, 0xB27C0020},
		{arm.X0, arm.X1, 255, 0xB2401C20},
	}

	for _, pattern := range usagePatterns {
		t.Logf("orr %s, %s, %d", pattern.Destination, pattern.Source, pattern.Number)
		code, encodable := arm.OrRegisterNumber(pattern.Destination, pattern.Source, pattern.Number)
		assert.True(t, encodable)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestOrRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Operand     cpu.Register
		Code        uint32
	}{
		{arm.X0, arm.X1, arm.X2, 0xAA020020},
	}

	for _, pattern := range usagePatterns {
		t.Logf("orr %s, %s, %s", pattern.Destination, pattern.Source, pattern.Operand)
		code := arm.OrRegisterRegister(pattern.Destination, pattern.Source, pattern.Operand)
		assert.Equal(t, code, pattern.Code)
	}
}