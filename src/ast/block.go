package ast

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// block returns the start and end of the block, the block's body, and an error if one occurred.
func block(tokens token.List, file *fs.File) (int, int, AST, error) {
	blockStart := tokens.IndexKind(token.BlockStart)
	blockEnd := tokens.LastIndexKind(token.BlockEnd)

	if blockStart == -1 {
		return 0, 0, nil, errors.New(MissingBlockStart, file, tokens[0].End())
	}

	if blockEnd == -1 {
		return 0, 0, nil, errors.New(MissingBlockEnd, file, tokens[len(tokens)-1].End())
	}

	body, err := Parse(tokens[blockStart+1:blockEnd], file)
	return blockStart, blockEnd, body, err
}