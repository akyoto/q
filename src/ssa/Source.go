package ssa

import "git.urbach.dev/cli/q/src/token"

type Source token.List

func (v Source) End() token.Position {
	return v[len(v)-1].End()
}

func (v *Source) SetSource(source token.List) {
	*v = Source(source)
}

func (v Source) Start() token.Position {
	return v[0].Position
}