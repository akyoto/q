package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestAndRegisterNumber(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Number      int
		Code        uint32
	}{
		{arm.X0, arm.X1, 1, 0x92400020},
		{arm.X0, arm.X1, 2, 0x927F0020},
		{arm.X0, arm.X1, 3, 0x92400420},
		{arm.X0, arm.X1, 7, 0x92400820},
		{arm.X0, arm.X1, 16, 0x927C0020},
		{arm.X0, arm.X1, 255, 0x92401C20},
	}

	for _, pattern := range usagePatterns {
		t.Logf("and %s, %s, %d", pattern.Destination, pattern.Source, pattern.Number)
		code, encodable := arm.AndRegisterNumber(pattern.Destination, pattern.Source, pattern.Number)
		assert.True(t, encodable)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestAndRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Operand     cpu.Register
		Code        uint32
	}{
		{arm.X0, arm.X1, arm.X2, 0x8A020020},
	}

	for _, pattern := range usagePatterns {
		t.Logf("and %s, %s, %s", pattern.Destination, pattern.Source, pattern.Operand)
		code := arm.AndRegisterRegister(pattern.Destination, pattern.Source, pattern.Operand)
		assert.Equal(t, code, pattern.Code)
	}
}