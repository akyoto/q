package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestLoadRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Base        cpu.Register
		Mode        arm.AddressMode
		Offset      int
		Length      byte
		Code        uint32
	}{
		{arm.X0, arm.X1, arm.UnscaledImmediate, -8, 1, 0x385F8020},
		{arm.X1, arm.X0, arm.UnscaledImmediate, -8, 1, 0x385F8001},
		{arm.X0, arm.X1, arm.UnscaledImmediate, -8, 2, 0x785F8020},
		{arm.X1, arm.X0, arm.UnscaledImmediate, -8, 2, 0x785F8001},
		{arm.X0, arm.X1, arm.UnscaledImmediate, -8, 4, 0xB85F8020},
		{arm.X1, arm.X0, arm.UnscaledImmediate, -8, 4, 0xB85F8001},
		{arm.X0, arm.X1, arm.UnscaledImmediate, -8, 8, 0xF85F8020},
		{arm.X1, arm.X0, arm.UnscaledImmediate, -8, 8, 0xF85F8001},
		{arm.X2, arm.X1, arm.UnscaledImmediate, -8, 8, 0xF85F8022},
		{arm.X2, arm.X1, arm.UnscaledImmediate, 0, 8, 0xF8400022},
		{arm.X2, arm.X1, arm.UnscaledImmediate, 8, 8, 0xF8408022},
		{arm.X2, arm.X1, arm.UnscaledImmediate, -256, 8, 0xF8500022},
		{arm.X2, arm.X1, arm.UnscaledImmediate, 255, 8, 0xF84FF022},

		{arm.X0, arm.SP, arm.PostIndex, 16, 8, 0xF84107E0},
		{arm.X1, arm.SP, arm.PostIndex, 16, 8, 0xF84107E1},
		{arm.X2, arm.SP, arm.PostIndex, 16, 8, 0xF84107E2},
	}

	for _, pattern := range usagePatterns {
		t.Logf("ldur %s, [%s, %d] %db", pattern.Destination, pattern.Base, pattern.Offset, pattern.Length)
		code := arm.LoadRegister(pattern.Destination, pattern.Base, pattern.Mode, pattern.Offset, pattern.Length)
		assert.Equal(t, code, pattern.Code)
	}
}