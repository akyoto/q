package ssa

import "git.urbach.dev/cli/q/src/token"

type HasSource interface {
	Start() token.Position
	End() token.Position
	StringFrom([]byte) string
}

// Source tracks the source tokens.
type Source token.Source

func (v Source) Start() token.Position         { return v.StartPos }
func (v Source) End() token.Position           { return v.EndPos }
func (v Source) StringFrom(code []byte) string { return string(code[v.Start():v.End()]) }