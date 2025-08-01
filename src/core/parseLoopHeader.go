package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
)

// parseLoopHeader parses the loop header.
func (f *Function) parseLoopHeader(head *expression.Expression) (name string, from *expression.Expression, to *expression.Expression) {
	switch head.Token.Kind {
	case token.Define:
		name = head.Children[0].String(f.File.Bytes)
		right := head.Children[1]
		from = right.Children[0]
		to = right.Children[1]
	case token.Range:
		name = fmt.Sprintf("_counter_%d", f.Count.Loop)
		from = head.Children[0]
		to = head.Children[1]
	}

	return
}