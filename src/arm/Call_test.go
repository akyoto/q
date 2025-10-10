package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestCall(t *testing.T) {
	usagePatterns := []struct {
		Offset int
		Code   uint32
	}{
		{0, 0x94000000},
		{1, 0x94000001},
		{-1, 0x97FFFFFF},
	}

	for _, pattern := range usagePatterns {
		t.Logf("bl %d", pattern.Offset)
		code, encodable := arm.Call(pattern.Offset)
		assert.True(t, encodable)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestCallRegister(t *testing.T) {
	usagePatterns := []struct {
		Register cpu.Register
		Code     uint32
	}{
		{arm.X0, 0xD63F0000},
		{arm.X1, 0xD63F0020},
		{arm.X2, 0xD63F0040},
	}

	for _, pattern := range usagePatterns {
		t.Logf("blr %d", pattern.Register)
		code := arm.CallRegister(pattern.Register)
		assert.Equal(t, code, pattern.Code)
	}
}