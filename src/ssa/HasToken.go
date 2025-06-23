package ssa

import "git.urbach.dev/cli/q/src/token"

type HasToken struct {
	Source token.Token
}

func (v *HasToken) Token() token.Token {
	return v.Source
}