package ast

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

func parseSwitch(tokens token.List, file *fs.File) (Node, error) {
	blockStart := tokens.IndexKind(token.BlockStart)
	blockEnd := tokens.LastIndexKind(token.BlockEnd)

	if blockStart == -1 {
		return nil, errors.NewAt(MissingBlockStart, file, tokens[0].End())
	}

	if blockEnd == -1 {
		return nil, errors.NewAt(MissingBlockEnd, file, tokens[len(tokens)-1].End())
	}

	body := tokens[blockStart+1 : blockEnd]

	if len(body) == 0 {
		return nil, errors.New(EmptySwitch, file, tokens[0])
	}

	cases, err := parseCases(body, file)
	return &Switch{Cases: cases}, err
}