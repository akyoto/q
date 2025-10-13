package scanner

import (
	"os"
	"path/filepath"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/global"
	"git.urbach.dev/cli/q/src/token"
)

// scanImport scans an import.
func (s *scanner) scanImport(file *fs.File, tokens token.List, i int) (int, error) {
	i++

	if tokens[i].Kind != token.Identifier {
		return i, errors.New(ExpectedPackageName, file, tokens[i])
	}

	packageName := tokens[i].StringFrom(file.Bytes)
	fullPath := filepath.Join(global.Library, packageName)
	stat, err := os.Stat(fullPath)

	if err != nil {
		return i, errors.New(&UnknownImport{Package: packageName}, file, tokens[i])
	}

	if !stat.IsDir() {
		return i, errors.New(&IsNotDirectory{Path: fullPath}, file, tokens[i])
	}

	file.Imports[packageName] = &fs.Import{
		Package: packageName,
		Tokens:  tokens[i : i+1],
	}

	s.queueDirectory(fullPath, packageName)
	return i, nil
}