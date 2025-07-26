package ast

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// block retrieves the start and end position of a block.
func block(tokens token.List, file *fs.File) (blockStart int, blockEnd int, body AST, err error) {
	blockStart = tokens.IndexKind(token.BlockStart)
	blockEnd = tokens.LastIndexKind(token.BlockEnd)

	if blockStart == -1 {
		err = errors.New(MissingBlockStart, file, tokens[0].End())
		return
	}

	if blockEnd == -1 {
		err = errors.New(MissingBlockEnd, file, tokens[len(tokens)-1].End())
		return
	}

	body, err = Parse(tokens[blockStart+1:blockEnd], file)
	return
}