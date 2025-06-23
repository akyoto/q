package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestStoreRegister(t *testing.T) {
	usagePatterns := []struct {
		Source cpu.Register
		Base   cpu.Register
		Offset int
		Length byte
		Code   uint32
	}{
		{arm.X0, arm.X1, -8, 1, 0x381F8020},
		{arm.X1, arm.X0, -8, 1, 0x381F8001},
		{arm.X0, arm.X1, -8, 2, 0x781F8020},
		{arm.X1, arm.X0, -8, 2, 0x781F8001},
		{arm.X0, arm.X1, -8, 4, 0xB81F8020},
		{arm.X1, arm.X0, -8, 4, 0xB81F8001},
		{arm.X0, arm.X1, -8, 8, 0xF81F8020},
		{arm.X1, arm.X0, -8, 8, 0xF81F8001},
	}

	for _, pattern := range usagePatterns {
		t.Logf("stur %s, [%s, #%d] %db", pattern.Source, pattern.Base, pattern.Offset, pattern.Length)
		code := arm.StoreRegister(pattern.Source, pattern.Base, pattern.Offset, pattern.Length)
		assert.Equal(t, code, pattern.Code)
	}
}