package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestXorRegisterNumber(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Number      int
		Code        uint32
	}{
		{arm.X0, arm.X1, 1, 0xD2400020},
		{arm.X0, arm.X1, 2, 0xD27F0020},
		{arm.X0, arm.X1, 3, 0xD2400420},
		{arm.X0, arm.X1, 7, 0xD2400820},
		{arm.X0, arm.X1, 16, 0xD27C0020},
		{arm.X0, arm.X1, 255, 0xD2401C20},
	}

	for _, pattern := range usagePatterns {
		t.Logf("eor %s, %s, %d", pattern.Destination, pattern.Source, pattern.Number)
		code, encodable := arm.XorRegisterNumber(pattern.Destination, pattern.Source, pattern.Number)
		assert.True(t, encodable)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestXorRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Operand     cpu.Register
		Code        uint32
	}{
		{arm.X0, arm.X1, arm.X2, 0xCA020020},
	}

	for _, pattern := range usagePatterns {
		t.Logf("eor %s, %s, %s", pattern.Destination, pattern.Source, pattern.Operand)
		code := arm.XorRegisterRegister(pattern.Destination, pattern.Source, pattern.Operand)
		assert.Equal(t, code, pattern.Code)
	}
}