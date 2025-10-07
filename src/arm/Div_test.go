package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestDivSignedRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Operand     cpu.Register
		Code        uint32
	}{
		{arm.X0, arm.X1, arm.X2, 0x9AC20C20},
	}

	for _, pattern := range usagePatterns {
		t.Logf("sdiv %s, %s, %s", pattern.Destination, pattern.Source, pattern.Operand)
		code := arm.DivSignedRegisterRegister(pattern.Destination, pattern.Source, pattern.Operand)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestDivUnsignedRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Operand     cpu.Register
		Code        uint32
	}{
		{arm.X0, arm.X1, arm.X2, 0x9AC20820},
	}

	for _, pattern := range usagePatterns {
		t.Logf("udiv %s, %s, %s", pattern.Destination, pattern.Source, pattern.Operand)
		code := arm.DivUnsignedRegisterRegister(pattern.Destination, pattern.Source, pattern.Operand)
		assert.Equal(t, code, pattern.Code)
	}
}