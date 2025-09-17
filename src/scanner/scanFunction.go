package scanner

import (
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// scanFunction scans a function.
func (s *scanner) scanFunction(file *fs.File, tokens token.List, i int) (int, error) {
	function, i, err := scanFunctionSignature(file, file.Package, tokens, i, token.BlockStart)

	if err != nil {
		return i, err
	}

	bodyStart, i, err := s.scanFunctionBody(file, tokens, i)

	if err != nil {
		return i, err
	}

	function.SetBody(bodyStart, i)
	s.functions <- function
	return i, nil
}