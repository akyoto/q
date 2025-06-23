package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestREX(t *testing.T) {
	testData := []struct{ w, r, x, b, expected byte }{
		{0, 0, 0, 0, 0b_0100_0000},
		{0, 0, 0, 1, 0b_0100_0001},
		{0, 0, 1, 0, 0b_0100_0010},
		{0, 0, 1, 1, 0b_0100_0011},
		{0, 1, 0, 0, 0b_0100_0100},
		{0, 1, 0, 1, 0b_0100_0101},
		{0, 1, 1, 0, 0b_0100_0110},
		{0, 1, 1, 1, 0b_0100_0111},
		{1, 0, 0, 0, 0b_0100_1000},
		{1, 0, 0, 1, 0b_0100_1001},
		{1, 0, 1, 0, 0b_0100_1010},
		{1, 0, 1, 1, 0b_0100_1011},
		{1, 1, 0, 0, 0b_0100_1100},
		{1, 1, 0, 1, 0b_0100_1101},
		{1, 1, 1, 0, 0b_0100_1110},
		{1, 1, 1, 1, 0b_0100_1111},
	}

	for _, test := range testData {
		rex := x86.REX(test.w, test.r, test.x, test.b)
		assert.Equal(t, rex, test.expected)
	}
}