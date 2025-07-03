package ssa

import "git.urbach.dev/cli/q/src/token"

type HasSource interface {
	Start() token.Position
	End() token.Position
}