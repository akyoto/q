package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestStoreDynamicRegister(t *testing.T) {
	usagePatterns := []struct {
		Source cpu.Register
		Base   cpu.Register
		Offset cpu.Register
		Scale  arm.ScaleFactor
		Length byte
		Code   uint32
	}{
		{arm.X0, arm.X1, arm.X2, arm.Scale1, 8, 0xF8226820},
		{arm.X0, arm.X1, arm.X2, arm.Scale1, 4, 0xB8226820},
		{arm.X0, arm.X1, arm.X2, arm.Scale1, 2, 0x78226820},
		{arm.X0, arm.X1, arm.X2, arm.Scale1, 1, 0x38226820},
	}

	for _, pattern := range usagePatterns {
		t.Logf("str %s, [%s, #%s] %db", pattern.Source, pattern.Base, pattern.Offset, pattern.Length)
		code := arm.StoreDynamicRegister(pattern.Source, pattern.Base, pattern.Offset, pattern.Scale, pattern.Length)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestStoreRegister(t *testing.T) {
	usagePatterns := []struct {
		Source cpu.Register
		Base   cpu.Register
		Mode   arm.AddressMode
		Offset int
		Length byte
		Code   uint32
	}{
		{arm.X0, arm.X1, arm.UnscaledImmediate, -8, 1, 0x381F8020},
		{arm.X1, arm.X0, arm.UnscaledImmediate, -8, 1, 0x381F8001},
		{arm.X0, arm.X1, arm.UnscaledImmediate, -8, 2, 0x781F8020},
		{arm.X1, arm.X0, arm.UnscaledImmediate, -8, 2, 0x781F8001},
		{arm.X0, arm.X1, arm.UnscaledImmediate, -8, 4, 0xB81F8020},
		{arm.X1, arm.X0, arm.UnscaledImmediate, -8, 4, 0xB81F8001},
		{arm.X0, arm.X1, arm.UnscaledImmediate, -8, 8, 0xF81F8020},
		{arm.X1, arm.X0, arm.UnscaledImmediate, -8, 8, 0xF81F8001},

		{arm.X0, arm.SP, arm.PreIndex, -16, 8, 0xF81F0FE0},
		{arm.X1, arm.SP, arm.PreIndex, -16, 8, 0xF81F0FE1},
		{arm.X2, arm.SP, arm.PreIndex, -16, 8, 0xF81F0FE2},
	}

	for _, pattern := range usagePatterns {
		t.Logf("stur %s, [%s, #%d] %db", pattern.Source, pattern.Base, pattern.Offset, pattern.Length)
		code := arm.StoreRegister(pattern.Source, pattern.Base, pattern.Mode, pattern.Offset, pattern.Length)
		assert.Equal(t, code, pattern.Code)
	}
}