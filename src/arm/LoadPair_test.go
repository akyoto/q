package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestLoadPair(t *testing.T) {
	usagePatterns := []struct {
		Reg1   cpu.Register
		Reg2   cpu.Register
		Base   cpu.Register
		Offset int
		Code   uint32
	}{
		{arm.FP, arm.LR, arm.SP, 32, 0xA8C27BFD},
		{arm.FP, arm.LR, arm.SP, 16, 0xA8C17BFD},
	}

	for _, pattern := range usagePatterns {
		t.Logf("ldp %s, %s, [%s], #%d", pattern.Reg1, pattern.Reg2, pattern.Base, pattern.Offset)
		code := arm.LoadPair(pattern.Reg1, pattern.Reg2, pattern.Base, pattern.Offset)
		assert.Equal(t, code, pattern.Code)
	}
}