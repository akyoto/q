package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestShiftLeftNumber(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Bits        int
		Code        uint32
	}{
		{arm.X0, arm.X0, 0, 0xD340FC00},
		{arm.X0, arm.X0, 1, 0xD37FF800},
		{arm.X0, arm.X0, 8, 0xD378DC00},
		{arm.X0, arm.X0, 16, 0xD370BC00},
		{arm.X0, arm.X0, 63, 0xD3410000},
	}

	for _, pattern := range usagePatterns {
		t.Logf("%b", pattern.Code)
		t.Logf("lsl %s, %s, %x", pattern.Destination, pattern.Source, pattern.Bits)
		code := arm.ShiftLeftNumber(pattern.Destination, pattern.Source, pattern.Bits)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestShiftRightSignedNumber(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Bits        int
		Code        uint32
	}{
		{arm.X0, arm.X0, 0, 0x9340FC00},
		{arm.X0, arm.X0, 1, 0x9341FC00},
		{arm.X0, arm.X0, 8, 0x9348FC00},
		{arm.X0, arm.X0, 16, 0x9350FC00},
		{arm.X0, arm.X0, 63, 0x937FFC00},
	}

	for _, pattern := range usagePatterns {
		t.Logf("asr %s, %s, %x", pattern.Destination, pattern.Source, pattern.Bits)
		code := arm.ShiftRightSignedNumber(pattern.Destination, pattern.Source, pattern.Bits)
		assert.Equal(t, code, pattern.Code)
	}
}