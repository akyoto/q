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
		return i, errors.NewAt(ExpectedPackageName, file, tokens[i].Position)
	}

	packageName := tokens[i].StringFrom(file.Bytes)
	fullPath := filepath.Join(global.Library, packageName)
	stat, err := os.Stat(fullPath)

	if err != nil {
		return i, errors.NewAt(&UnknownImport{Package: packageName}, file, tokens[i].Position)
	}

	if !stat.IsDir() {
		return i, errors.NewAt(&IsNotDirectory{Path: fullPath}, file, tokens[i].Position)
	}

	file.Imports[packageName] = &fs.Import{
		Package:  packageName,
		Position: tokens[i].Position,
	}

	s.queueDirectory(fullPath, packageName)
	return i, nil
}