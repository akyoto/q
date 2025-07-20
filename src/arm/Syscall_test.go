package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/go/assert"
)

func TestSyscall(t *testing.T) {
	usagePatterns := []struct {
		Number int
		Code   uint32
	}{
		{0, 0xD4000001},
		{0xFFFF, 0xD41FFFE1},
	}

	for _, pattern := range usagePatterns {
		t.Logf("svc %d", pattern.Number)
		code := arm.Syscall(pattern.Number)
		assert.Equal(t, code, pattern.Code)
	}
}