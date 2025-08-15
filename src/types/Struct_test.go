package types_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestStruct(t *testing.T) {
	point := types.NewStruct(nil, "math", "Point")

	x := &types.Field{
		Name:   "x",
		Type:   types.Int,
		Index:  0,
		Offset: 0,
	}

	y := &types.Field{
		Name:   "y",
		Type:   types.Int,
		Index:  1,
		Offset: 8,
	}

	point.AddField(x)
	point.AddField(y)

	assert.Equal(t, x.String(), "x")
	assert.Equal(t, y.String(), "y")
	assert.Equal(t, point.FieldByName("x"), x)
	assert.Equal(t, point.FieldByName("y"), y)
	assert.Nil(t, point.FieldByName("invalid"))
}