package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestLoadTLS(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Code        uint32
	}{
		{arm.X0, 0xD53BD040},
		{arm.X1, 0xD53BD041},
	}

	for _, pattern := range usagePatterns {
		t.Logf("mrs %s, TPIDR_EL0", pattern.Destination)
		code := arm.LoadTLS(pattern.Destination)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestStoreTLS(t *testing.T) {
	usagePatterns := []struct {
		Source cpu.Register
		Code   uint32
	}{
		{arm.X0, 0xD51BD040},
		{arm.X1, 0xD51BD041},
	}

	for _, pattern := range usagePatterns {
		t.Logf("msr TPIDR_EL0, %s", pattern.Source)
		code := arm.StoreTLS(pattern.Source)
		assert.Equal(t, code, pattern.Code)
	}
}