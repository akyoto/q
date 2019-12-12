package build

import (
	"os"
	"strings"
	"sync"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// Scan scans the input file.
func (file *File) Scan(imports chan<- *Import, functions chan<- *Function) error {
	var (
		tokens                  = file.tokens
		newlines                = 0
		index    token.Position = 0
		t        token.Token
	)

begin:
	for ; index < len(tokens); index++ {
		t = tokens[index]

		if t.Kind != token.NewLine {
			newlines = 0
		}

		switch t.Kind {
		case token.Identifier:
			var function *Function
			var err error
			function, index, err = file.scanFunction(tokens, index)

			if err != nil {
				return err
			}

			functions <- function

		case token.Keyword:
			if t.Text() == "import" {
				var imp *Import
				var err error

				imp, index, err = file.scanImport(tokens, index)

				if err != nil {
					return err
				}

				file.imports[imp.BaseName] = imp
				imports <- imp
				goto begin
			}

			return NewError(errors.TopLevel, file.path, tokens[:index+1], nil)

		case token.NewLine:
			newlines++

			if newlines == 3 {
				return NewError(errors.UnnecessaryNewlines, file.path, tokens[:index+1], nil)
			}

		case token.Comment:
			// OK.

		default:
			return NewError(errors.TopLevel, file.path, tokens[:index+1], nil)
		}
	}

	return nil
}

// scanFunction scans a function.
func (file *File) scanFunction(tokens token.List, index token.Position) (*Function, token.Position, error) {
	var (
		groupLevel = 0
		blockLevel = 0
		newlines   = 0
	)

	functionName := tokens[index].Text()

	if functionName == "func" || functionName == "fn" {
		return nil, index, NewError(errors.InvalidFunctionName, file.path, tokens[:index+1], nil)
	}

	if index+1 >= len(tokens) || tokens[index+1].Kind != token.GroupStart {
		return nil, index, NewError(errors.ParameterOpeningBracket, file.path, tokens[:index+2], nil)
	}

	function := &Function{
		Name:           functionName,
		File:           file,
		parameterStart: index + 2,
	}

	function.Finished = sync.NewCond(&function.FinishedMutex)

	if functionName == "main" {
		function.CallCount = 1
	}

	file.functionCount++
	index++

	for ; index < len(tokens); index++ {
		t := tokens[index]

		if t.Kind != token.NewLine {
			newlines = 0
		}

		switch t.Kind {
		case token.BlockStart:
			if groupLevel > 0 {
				return function, index, NewError(&errors.MissingCharacter{Character: ")"}, file.path, tokens[:index+1], function)
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
			return function, index, nil

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
					return function, index, NewError(&errors.MissingType{Of: parameterName.Text()}, file.path, tokens[:function.parameterStart+1], function)
				}

				typeName := parameter[1].Text()
				typ := file.environment.Types[typeName]

				if typ == nil {
					return function, index, NewError(&errors.UnknownType{Name: typeName}, file.path, tokens[:index], function)
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
				return function, index, NewError(&errors.MissingType{Of: parameterName.Text()}, file.path, tokens[:function.parameterStart+1], function)
			}

			typeName := parameter[1].Text()
			typ := file.environment.Types[typeName]

			if typ == nil {
				return function, index, NewError(&errors.UnknownType{Name: typeName}, file.path, tokens[:index], function)
			}

			function.Parameters = append(function.Parameters, &Parameter{
				Name:     parameterName.Text(),
				Type:     typ,
				Position: function.parameterStart,
			})

			function.parameterStart = index + 1

		case token.Operator:
			if groupLevel != 0 || function == nil || function.TokenStart != 0 || t.Text() != "->" {
				continue
			}

			// Return type
			index++
			t = tokens[index]

			if t.Kind != token.Identifier {
				return function, index, NewError(errors.MissingReturnType, file.path, tokens[:index+1], function)
			}

			typeName := t.Text()
			typ := file.environment.Types[typeName]

			if typ == nil {
				return function, index, NewError(&errors.UnknownType{Name: typeName}, file.path, tokens[:index+1], function)
			}

			function.ReturnTypes = append(function.ReturnTypes, typ)

		case token.NewLine:
			newlines++

			if newlines == 3 {
				return function, index, NewError(errors.UnnecessaryNewlines, file.path, tokens[:index+1], function)
			}
		}
	}

	return function, index, nil
}

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
				return nil, index, NewError(&errors.InvalidCharacter{Character: t.Text()}, file.path, tokens[:index+1], nil)
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
				return nil, index, NewError(&errors.ImportNameAlreadyExists{ImportPath: otherImport.Path, Name: baseName}, file.path, tokens[:index], nil)
			}

			stat, err := os.Stat(imp.FullPath)

			if err != nil || !stat.IsDir() {
				return nil, index, NewError(&errors.PackageDoesntExist{ImportPath: imp.Path, FilePath: imp.FullPath}, file.path, tokens[:imp.Position+2], nil)
			}

			index++
			return imp, index, nil
		}
	}

	return nil, index, errors.InvalidExpression
}
