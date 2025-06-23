package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestStorePair(t *testing.T) {
	usagePatterns := []struct {
		Reg1   cpu.Register
		Reg2   cpu.Register
		Base   cpu.Register
		Offset int
		Code   uint32
	}{
		{arm.FP, arm.LR, arm.SP, -32, 0xA9BE7BFD},
		{arm.FP, arm.LR, arm.SP, -16, 0xA9BF7BFD},
	}

	for _, pattern := range usagePatterns {
		t.Logf("stp %s, %s, [%s, #%d]!", pattern.Reg1, pattern.Reg2, pattern.Base, pattern.Offset)
		code := arm.StorePair(pattern.Reg1, pattern.Reg2, pattern.Base, pattern.Offset)
		assert.Equal(t, code, pattern.Code)
	}
}