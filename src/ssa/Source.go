package ssa

import "git.urbach.dev/cli/q/src/token"

// HasSource is an interface for values that have a source position.
type HasSource interface {
	// Start returns the start position of the value.
	Start() token.Position

	// End returns the end position of the value.
	End() token.Position

	// StringFrom returns the source code of the value.
	StringFrom([]byte) string
}

// Source tracks the source tokens.
type Source = token.Source