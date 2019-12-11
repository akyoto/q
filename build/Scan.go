package build

import (
	"os"
	"strings"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// Scan scans the input file.
func (file *File) Scan(imports chan<- *Import, functions chan<- *Function) error {
	var (
		function   *Function
		groupLevel                = 0
		blockLevel                = 0
		tokens                    = file.tokens
		newlines                  = 0
		index      token.Position = 0
		t          token.Token
	)

begin:
	for ; index < len(tokens); index++ {
		t = tokens[index]

		if t.Kind != token.NewLine {
			newlines = 0
		}

		switch t.Kind {
		case token.Identifier:
			if function != nil {
				continue
			}

			functionName := t.Text()

			if functionName == "func" || functionName == "fn" {
				return NewError(errors.InvalidFunctionName, file.path, tokens[:index+1], function)
			}

			if index+1 >= len(tokens) || tokens[index+1].Kind != token.GroupStart {
				return NewError(errors.ParameterOpeningBracket, file.path, tokens[:index+2], function)
			}

			function = &Function{
				Name:           functionName,
				File:           file,
				Finished:       make(chan struct{}),
				parameterStart: index + 2,
			}

			if functionName == "main" {
				function.CallCount = 1
			}

			file.functionCount++

		case token.BlockStart:
			if groupLevel > 0 {
				return NewError(&errors.MissingCharacter{Character: ")"}, file.path, tokens[:index+1], function)
			}

			blockLevel++

			if function.TokenStart != 0 {
				continue
			}

			function.TokenStart = index + 1

		case token.BlockEnd:
			blockLevel--

			if blockLevel != 0 {
				continue
			}

			function.TokenEnd = index
			function.Name = PolymorphName(function.Name, len(function.Parameters))
			functions <- function
			function = nil

		case token.GroupStart:
			groupLevel++

		case token.GroupEnd:
			groupLevel--

			if groupLevel != 0 {
				continue
			}

			if function.TokenStart != 0 {
				continue
			}

			if function.parameterStart < index {
				parameter := tokens[function.parameterStart:index]
				parameterName := parameter[0]

				if len(parameter) == 1 {
					return NewError(&errors.MissingType{Of: parameterName.Text()}, file.path, tokens[:function.parameterStart+1], function)
				}

				typeName := parameter[1].Text()
				typ := file.environment.Types[typeName]

				if typ == nil {
					return NewError(&errors.UnknownType{Name: typeName}, file.path, tokens[:index], function)
				}

				function.Parameters = append(function.Parameters, &Parameter{
					Name:     parameterName.Text(),
					Type:     typ,
					Position: function.parameterStart,
				})

				function.parameterStart = -1
			}

		case token.Separator:
			if function == nil || function.TokenStart != 0 || groupLevel != 1 {
				continue
			}

			if function.parameterStart >= index {
				continue
			}

			parameter := tokens[function.parameterStart:index]
			parameterName := parameter[0]

			if len(parameter) == 1 {
				return NewError(&errors.MissingType{Of: parameterName.Text()}, file.path, tokens[:function.parameterStart+1], function)
			}

			typeName := parameter[1].Text()
			typ := file.environment.Types[typeName]

			if typ == nil {
				return NewError(&errors.UnknownType{Name: typeName}, file.path, tokens[:index], function)
			}

			function.Parameters = append(function.Parameters, &Parameter{
				Name:     parameterName.Text(),
				Type:     typ,
				Position: function.parameterStart,
			})

			function.parameterStart = index + 1

		case token.Keyword:
			if function != nil {
				continue
			}

			if t.Text() != "import" {
				return NewError(errors.TopLevel, file.path, tokens[:index+1], function)
			}

			var baseName string
			position := index
			fullImportPath := strings.Builder{}
			fullImportPath.WriteString(file.environment.StandardLibrary)
			fullImportPath.WriteByte('/')
			importPath := strings.Builder{}
			index++

			for ; index < len(tokens); index++ {
				t = tokens[index]

				switch t.Kind {
				case token.Identifier:
					baseName = t.Text()
					fullImportPath.WriteString(baseName)
					importPath.WriteString(baseName)

				case token.Operator:
					if t.Text() != "." {
						return NewError(&errors.InvalidCharacter{Character: t.Text()}, file.path, tokens[:index+1], function)
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
						return NewError(&errors.ImportNameAlreadyExists{ImportPath: otherImport.Path, Name: baseName}, file.path, tokens[:index], function)
					}

					stat, err := os.Stat(imp.FullPath)

					if err != nil || !stat.IsDir() {
						return NewError(&errors.PackageDoesntExist{ImportPath: imp.Path, FilePath: imp.FullPath}, file.path, tokens[:imp.Position+2], function)
					}

					file.imports[baseName] = imp
					imports <- imp
					index++
					goto begin
				}
			}

		case token.Operator:
			if groupLevel != 0 || function == nil || function.TokenStart != 0 || t.Text() != "->" {
				continue
			}

			// Return type
			index++
			t = tokens[index]

			if t.Kind != token.Identifier {
				return NewError(errors.MissingReturnType, file.path, tokens[:index+1], function)
			}

			typeName := t.Text()
			typ := file.environment.Types[typeName]

			if typ == nil {
				return NewError(&errors.UnknownType{Name: typeName}, file.path, tokens[:index+1], function)
			}

			function.ReturnTypes = append(function.ReturnTypes, typ)

		case token.NewLine:
			newlines++

			if newlines == 3 {
				return NewError(errors.UnnecessaryNewlines, file.path, tokens[:index+1], function)
			}

		case token.Comment:
			// OK.

		default:
			if function == nil {
				return NewError(errors.TopLevel, file.path, tokens[:index+1], function)
			}
		}
	}

	return nil
}
