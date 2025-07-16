package token_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestTokenGroups(t *testing.T) {
	assert.True(t, token.Assign.IsAssignment())
	assert.True(t, token.Add.IsOperator())
	assert.True(t, token.If.IsKeyword())
	assert.True(t, token.Not.IsUnaryOperator())
	assert.True(t, token.Number.IsNumeric())
	assert.True(t, token.Equal.IsComparison())
}