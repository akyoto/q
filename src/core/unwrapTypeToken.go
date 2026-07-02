package core

import (
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// unwrapTypeToken turns a single `Type` kind token into the original list of tokens.
func unwrapTypeToken(typ token.Token, file *fs.File) token.List {
	start := 0
	end := 0

	for i, t := range file.Tokens {
		if t.Position == typ.Position {
			start = i
		}

		if t.End() == typ.End() {
			end = i + 1
			break
		}
	}

	return file.Tokens[start:end]
}