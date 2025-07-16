package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/token"
)

// block retrieves the start and end position of a block.
func (f *Function) block(tokens token.List) (blockStart int, blockEnd int, err error) {
	blockStart = tokens.IndexKind(token.BlockStart)

	if blockStart == -1 {
		err = errors.New(nil, f.File, tokens[0].End())
		return
	}

	blockEnd = tokens.LastIndexKind(token.BlockEnd)

	if blockEnd == -1 {
		err = errors.New(nil, f.File, tokens[len(tokens)-1].End())
		return
	}

	return
}