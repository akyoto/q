package sizeof_test

import (
	"math"
	"testing"

	"git.urbach.dev/cli/q/src/sizeof"
	"git.urbach.dev/go/assert"
)

func TestSigned(t *testing.T) {
	assert.Equal(t, sizeof.Signed(0), 1)
	assert.Equal(t, sizeof.Signed(math.MinInt8), 1)
	assert.Equal(t, sizeof.Signed(math.MaxInt8), 1)
	assert.Equal(t, sizeof.Signed(math.MinInt16), 2)
	assert.Equal(t, sizeof.Signed(math.MaxInt16), 2)
	assert.Equal(t, sizeof.Signed(math.MinInt32), 4)
	assert.Equal(t, sizeof.Signed(math.MaxInt32), 4)
	assert.Equal(t, sizeof.Signed(math.MinInt64), 8)
	assert.Equal(t, sizeof.Signed(math.MaxInt64), 8)
}