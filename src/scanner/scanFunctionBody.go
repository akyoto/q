package scanner

import (
	"path/filepath"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/global"
	"git.urbach.dev/cli/q/src/token"
)

// scanFunctionBody scans the body of a function.
func (s *scanner) scanFunctionBody(file *fs.File, tokens token.List, i int) (int, int, error) {
	var (
		blockLevel = 0
		bodyStart  = -1
	)

	for i < len(tokens) {
		switch tokens[i].Kind {
		case token.BlockStart:
			blockLevel++
			i++

			if blockLevel == 1 {
				bodyStart = i
			}

			continue

		case token.BlockEnd:
			blockLevel--

			if blockLevel < 0 {
				return bodyStart, i, errors.NewAt(MissingBlockStart, file, tokens[i].Position)
			}

			if blockLevel == 0 {
				return bodyStart, i, nil
			}

			i++
			continue

		case token.Go:
			s.queueDirectory(filepath.Join(global.Library, "thread"), "thread")

		case token.New, token.Delete:
			s.queueDirectory(filepath.Join(global.Library, "mem"), "mem")

		case token.Invalid:
			return bodyStart, i, errors.New(&InvalidCharacter{Character: tokens[i].StringFrom(file.Bytes)}, file, tokens[i])

		case token.EOF:
			if blockLevel > 0 {
				return bodyStart, i, errors.NewAt(MissingBlockEnd, file, tokens[i].Position)
			}

			return bodyStart, i, errors.NewAt(ExpectedFunctionDefinition, file, tokens[i].Position)
		}

		if blockLevel > 0 {
			i++
			continue
		}

		return bodyStart, i, errors.NewAt(ExpectedFunctionDefinition, file, tokens[i].Position)
	}

	panic("no EOF token")
}