package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestSet(t *testing.T) {
	usagePatterns := []struct {
		Register  cpu.Register
		Condition int
		Code      []byte
	}{
		{x86.R0, 1, []byte{0x0F, 0x94, 0xC0}},
		{x86.R1, 1, []byte{0x0F, 0x94, 0xC1}},
		{x86.R0, 2, []byte{0x0F, 0x95, 0xC0}},
		{x86.R1, 2, []byte{0x0F, 0x95, 0xC1}},
		{x86.R0, 3, []byte{0x0F, 0x9F, 0xC0}},
		{x86.R1, 3, []byte{0x0F, 0x9F, 0xC1}},
		{x86.R0, 4, []byte{0x0F, 0x9D, 0xC0}},
		{x86.R1, 4, []byte{0x0F, 0x9D, 0xC1}},
		{x86.R0, 5, []byte{0x0F, 0x9C, 0xC0}},
		{x86.R1, 5, []byte{0x0F, 0x9C, 0xC1}},
		{x86.R0, 6, []byte{0x0F, 0x9E, 0xC0}},
		{x86.R1, 6, []byte{0x0F, 0x9E, 0xC1}},
	}

	for _, pattern := range usagePatterns {
		var code []byte

		switch pattern.Condition {
		case 1:
			t.Logf("sete %s", pattern.Register)
			code = x86.SetIfEqual(nil, pattern.Register)
		case 2:
			t.Logf("setne %s", pattern.Register)
			code = x86.SetIfNotEqual(nil, pattern.Register)
		case 3:
			t.Logf("setg %s", pattern.Register)
			code = x86.SetIfGreater(nil, pattern.Register)
		case 4:
			t.Logf("setge %s", pattern.Register)
			code = x86.SetIfGreaterOrEqual(nil, pattern.Register)
		case 5:
			t.Logf("setl %s", pattern.Register)
			code = x86.SetIfLess(nil, pattern.Register)
		case 6:
			t.Logf("setle %s", pattern.Register)
			code = x86.SetIfLessOrEqual(nil, pattern.Register)
		}

		assert.DeepEqual(t, code, pattern.Code)
	}
}