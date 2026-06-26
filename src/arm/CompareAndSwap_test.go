package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestCompareAndSwap(t *testing.T) {
	usagePatterns := []struct {
		OldValue cpu.Register
		NewValue cpu.Register
		Address  cpu.Register
		Length   byte
		Code     uint32
	}{
		{arm.X0, arm.X1, arm.X2, 4, 0x88E0FC41},
		{arm.X2, arm.X1, arm.X0, 4, 0x88E2FC01},
	}

	for _, pattern := range usagePatterns {
		t.Logf("casal %s, %s, [#%s] %db", pattern.OldValue, pattern.NewValue, pattern.Address, pattern.Length)
		code := arm.CompareAndSwap(pattern.OldValue, pattern.NewValue, pattern.Address, pattern.Length)
		assert.Equal(t, code, pattern.Code)
	}
}