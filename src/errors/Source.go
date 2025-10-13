package errors

import "git.urbach.dev/cli/q/src/token"

// Source is an interface for values that have a source region.
type Source interface {
	End() token.Position
	Start() token.Position
	StringFrom([]byte) string
}