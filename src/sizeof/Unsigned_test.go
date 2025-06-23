package sizeof_test

import (
	"math"
	"testing"

	"git.urbach.dev/cli/q/src/sizeof"
	"git.urbach.dev/go/assert"
)

func TestUnsigned(t *testing.T) {
	assert.Equal(t, sizeof.Unsigned(0), 1)
	assert.Equal(t, sizeof.Unsigned(math.MaxUint8), 1)
	assert.Equal(t, sizeof.Unsigned(math.MaxUint8+1), 2)
	assert.Equal(t, sizeof.Unsigned(math.MaxUint16), 2)
	assert.Equal(t, sizeof.Unsigned(math.MaxUint16+1), 4)
	assert.Equal(t, sizeof.Unsigned(math.MaxUint32), 4)
	assert.Equal(t, sizeof.Unsigned(math.MaxUint32+1), 8)
}