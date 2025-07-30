package scanner

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// scanFunction scans a function.
func (s *scanner) scanFunction(file *fs.File, tokens token.List, i int) (int, error) {
	function, i, err := scanSignature(file, file.Package, tokens, i, token.BlockStart)

	if err != nil {
		return i, err
	}

	var (
		blockLevel = 0
		bodyStart  = -1
	)

	// Function definition
	for i < len(tokens) {
		if tokens[i].Kind == token.BlockStart {
			blockLevel++
			i++

			if blockLevel == 1 {
				bodyStart = i
			}

			continue
		}

		if tokens[i].Kind == token.BlockEnd {
			blockLevel--

			if blockLevel < 0 {
				return i, errors.New(MissingBlockStart, file, tokens[i].Position)
			}

			if blockLevel == 0 {
				break
			}

			i++
			continue
		}

		if tokens[i].Kind == token.Invalid {
			return i, errors.New(&InvalidCharacter{Character: tokens[i].String(file.Bytes)}, file, tokens[i].Position)
		}

		if tokens[i].Kind == token.EOF {
			if blockLevel > 0 {
				return i, errors.New(MissingBlockEnd, file, tokens[i].Position)
			}

			return i, errors.New(ExpectedFunctionDefinition, file, tokens[i].Position)
		}

		if blockLevel > 0 {
			i++
			continue
		}

		return i, errors.New(ExpectedFunctionDefinition, file, tokens[i].Position)
	}

	function.Body = tokens[bodyStart:i]
	s.functions <- function
	return i, nil
}