package token_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestSource(t *testing.T) {
	data := []byte("123")
	source := token.NewSource(0, 1)
	assert.Equal(t, source.StringFrom(data), "1")
	source.SetStart(1)
	source.SetEnd(2)
	assert.Equal(t, source.StringFrom(data), "2")
}