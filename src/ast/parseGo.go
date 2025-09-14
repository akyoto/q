package ast

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

func parseGo(tokens token.List, file *fs.File) (Node, error) {
	call := expression.Parse(tokens[1:])
	return &Go{Call: call}, nil
}