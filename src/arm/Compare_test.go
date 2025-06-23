package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestCompareRegisterNumber(t *testing.T) {
	usagePatterns := []struct {
		Source cpu.Register
		Number int
		Code   uint32
	}{
		{arm.X0, 0, 0xF100001F},
		{arm.X0, 1, 0xF100041F},
		{arm.X0, -1, 0xB100041F},
		{arm.X0, 0x1000, 0xF140041F},
	}

	for _, pattern := range usagePatterns {
		t.Logf("cmp %s, %d", pattern.Source, pattern.Number)
		code, encodable := arm.CompareRegisterNumber(pattern.Source, pattern.Number)
		assert.True(t, encodable)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestCompareRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Left  cpu.Register
		Right cpu.Register
		Code  uint32
	}{
		{arm.X0, arm.X1, 0xEB01001F},
	}

	for _, pattern := range usagePatterns {
		t.Logf("cmp %s, %s", pattern.Left, pattern.Right)
		code := arm.CompareRegisterRegister(pattern.Left, pattern.Right)
		assert.Equal(t, code, pattern.Code)
	}
}