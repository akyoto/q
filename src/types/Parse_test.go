package types_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		Source       string
		ExpectedType types.Type
	}{
		{"int", types.Int},
		{"int64", types.Int64},
		{"int32", types.Int32},
		{"int16", types.Int16},
		{"int8", types.Int8},
		{"uint", types.UInt},
		{"uint64", types.UInt64},
		{"uint32", types.UInt32},
		{"uint16", types.UInt16},
		{"uint8", types.UInt8},
		{"byte", types.Byte},
		{"bool", types.Bool},
		{"float", types.Float},
		{"float64", types.Float64},
		{"float32", types.Float32},
		{"any", types.Any},
		{"*any", types.AnyPointer},
		{"*byte", &types.Pointer{To: types.Byte}},
		{"[]any", types.AnyArray},
		{"[]byte", &types.Array{Of: types.Byte}},
		{"123", nil},
		{"*", nil},
		{"[]", nil},
		{"[", nil},
		{"_", nil},
	}

	for _, test := range tests {
		t.Run(test.Source, func(t *testing.T) {
			src := []byte(test.Source)
			tokens := token.Tokenize(src)
			typ := types.Parse(tokens, src)
			assert.True(t, types.Is(typ, test.ExpectedType))
		})
	}
}

func TestParseNil(t *testing.T) {
	tokens := []token.Token(nil)
	assert.Nil(t, types.Parse(tokens, nil))
}