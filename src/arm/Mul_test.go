package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestMulRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Operand     cpu.Register
		Code        uint32
	}{
		{arm.X0, arm.X1, arm.X2, 0x9B027C20},
	}

	for _, pattern := range usagePatterns {
		t.Logf("mul %s, %s, %s", pattern.Destination, pattern.Source, pattern.Operand)
		code := arm.MulRegisterRegister(pattern.Destination, pattern.Source, pattern.Operand)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestMultiplySubtract(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Operand     cpu.Register
		Extra       cpu.Register
		Code        uint32
	}{
		{arm.X0, arm.X1, arm.X2, arm.X3, 0x9B028C20},
		{arm.X3, arm.X0, arm.X2, arm.X1, 0x9B028403},
	}

	for _, pattern := range usagePatterns {
		t.Logf("msub %s, %s, %s, %s", pattern.Destination, pattern.Source, pattern.Operand, pattern.Extra)
		code := arm.MultiplySubtract(pattern.Destination, pattern.Source, pattern.Operand, pattern.Extra)
		assert.Equal(t, code, pattern.Code)
	}
}