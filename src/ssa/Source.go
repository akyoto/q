package ssa

import "git.urbach.dev/cli/q/src/token"

type Source token.List

func (v Source) Start() token.Position {
	return v[0].Position
}

func (v Source) End() token.Position {
	return v[len(v)-1].End()
}