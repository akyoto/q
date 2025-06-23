package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
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
		code := arm.Call(pattern.Offset)
		assert.Equal(t, code, pattern.Code)
	}
}