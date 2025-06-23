package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestSIB(t *testing.T) {
	testData := []struct{ scale, index, base, expected byte }{
		{0b_00, 0b_111, 0b_000, 0b_00_111_000},
		{0b_00, 0b_110, 0b_001, 0b_00_110_001},
		{0b_00, 0b_101, 0b_010, 0b_00_101_010},
		{0b_00, 0b_100, 0b_011, 0b_00_100_011},
		{0b_00, 0b_011, 0b_100, 0b_00_011_100},
		{0b_00, 0b_010, 0b_101, 0b_00_010_101},
		{0b_00, 0b_001, 0b_110, 0b_00_001_110},
		{0b_00, 0b_000, 0b_111, 0b_00_000_111},
		{0b_11, 0b_111, 0b_000, 0b_11_111_000},
		{0b_11, 0b_110, 0b_001, 0b_11_110_001},
		{0b_11, 0b_101, 0b_010, 0b_11_101_010},
		{0b_11, 0b_100, 0b_011, 0b_11_100_011},
		{0b_11, 0b_011, 0b_100, 0b_11_011_100},
		{0b_11, 0b_010, 0b_101, 0b_11_010_101},
		{0b_11, 0b_001, 0b_110, 0b_11_001_110},
		{0b_11, 0b_000, 0b_111, 0b_11_000_111},
	}

	for _, test := range testData {
		sib := x86.SIB(test.scale, test.index, test.base)
		assert.Equal(t, sib, test.expected)
	}
}