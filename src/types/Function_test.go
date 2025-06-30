package types_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestFunction(t *testing.T) {
	f := &types.Function{}
	assert.Equal(t, f.Name(), "()")
	assert.Equal(t, f.Size(), 8)

	f.Input = append(f.Input, types.Int64)
	assert.Equal(t, f.Name(), "(int64)")

	f.Input = append(f.Input, types.Int64)
	assert.Equal(t, f.Name(), "(int64, int64)")

	f.Output = append(f.Output, types.Int64)
	assert.Equal(t, f.Name(), "(int64, int64) -> (int64)")

	f.Output = append(f.Output, types.Int64)
	assert.Equal(t, f.Name(), "(int64, int64) -> (int64, int64)")
}