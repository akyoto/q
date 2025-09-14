package scanner

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// scanExtern scans a block of external libraries.
func (s *scanner) scanExtern(file *fs.File, tokens token.List, i int) (int, error) {
	i++

	if tokens[i].Kind != token.BlockStart {
		return i, errors.New(MissingBlockStart, file, tokens[i].Position)
	}

	for i < len(tokens) {
		switch tokens[i].Kind {
		case token.Identifier:
			var err error
			i, err = s.scanExternLibrary(file, tokens, i)

			if err != nil {
				return i, err
			}

		case token.BlockEnd:
			return i, nil
		}

		i++
	}

	return i, errors.New(MissingBlockEnd, file, tokens[i].Position)
}

// scanExternLibrary scans a block of external function declarations.
func (s *scanner) scanExternLibrary(file *fs.File, tokens token.List, i int) (int, error) {
	dllName := tokens[i].String(file.Bytes)
	i++

	if tokens[i].Kind != token.BlockStart {
		return i, errors.New(MissingBlockStart, file, tokens[i].Position)
	}

	i++

	for i < len(tokens) {
		switch tokens[i].Kind {
		case token.Identifier:
			function, j, err := scanFunctionSignature(file, dllName, tokens, i, token.NewLine)

			if err != nil {
				return j, err
			}

			i = j
			s.functions <- function

		case token.BlockEnd:
			return i, nil
		}

		i++
	}

	return i, errors.New(MissingBlockEnd, file, tokens[i].Position)
}