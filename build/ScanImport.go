package build

import (
	"os"
	"strings"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// scanImport scans an imports statement.
func (file *File) scanImport(tokens token.List, index token.Position) (*Import, token.Position, error) {
	var baseName string
	position := index
	fullImportPath := strings.Builder{}
	fullImportPath.WriteString(file.environment.StandardLibrary)
	fullImportPath.WriteByte('/')
	importPath := strings.Builder{}
	index++

	for ; index < len(tokens); index++ {
		t := tokens[index]

		switch t.Kind {
		case token.Identifier:
			baseName = t.Text()
			fullImportPath.WriteString(baseName)
			importPath.WriteString(baseName)

		case token.Operator:
			if t.Text() != "." {
				return nil, index, NewError(errors.New(&errors.InvalidCharacter{Character: t.Text()}), file.path, tokens[:index+1], nil)
			}

			fullImportPath.WriteByte('/')
			importPath.WriteByte('.')

		case token.NewLine:
			imp := &Import{
				Path:     importPath.String(),
				FullPath: fullImportPath.String(),
				BaseName: baseName,
				Position: position,
				Used:     0,
			}

			otherImport, exists := file.imports[baseName]

			if exists {
				return nil, index, NewError(errors.New(&errors.ImportNameAlreadyExists{ImportPath: otherImport.Path, Name: baseName}), file.path, tokens[:index], nil)
			}

			stat, err := os.Stat(imp.FullPath)

			if err != nil || !stat.IsDir() {
				return nil, index, NewError(errors.New(&errors.PackageDoesntExist{ImportPath: imp.Path, FilePath: imp.FullPath}), file.path, tokens[:imp.Position+2], nil)
			}

			index++
			return imp, index, nil
		}
	}

	return nil, index, errors.New(errors.InvalidExpression)
}
