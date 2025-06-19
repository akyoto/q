package scanner

import (
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// scanFunction scans a function.
func (s *scanner) scanFunction(file *fs.File, tokens token.List, i int) (int, error) {
	for i < len(tokens) && tokens[i].Kind != token.BlockEnd {
		i++
	}

	return i, nil
}