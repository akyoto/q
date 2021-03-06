package build

import (
	"sync"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// scanFunction scans a function.
func (file *File) scanFunction(tokens token.List, index token.Position) (*Function, token.Position, error) {
	var (
		groupLevel = 0
		blockLevel = 0
		newlines   = 0
	)

	functionName := tokens[index].Text()

	if functionName == "func" || functionName == "fn" {
		return nil, index, NewError(errors.New(errors.InvalidFunctionName), file.path, tokens[:index+1], nil)
	}

	if index+1 >= len(tokens) || tokens[index+1].Kind != token.GroupStart {
		return nil, index, NewError(errors.New(errors.ParameterOpeningBracket), file.path, tokens[:index+2], nil)
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
				return function, index, NewError(errors.New(&errors.MissingCharacter{Character: ")"}), file.path, tokens[:index+1], function)
			}

			blockLevel++

			if function.TokenStart != 0 {
				continue
			}

			function.TokenStart = index + 1

			// Return type
			if function.returnTypeStart != 0 {
				function.ReturnTypeTokens = tokens[function.returnTypeStart:index]

				if len(function.ReturnTypeTokens) == 0 {
					return function, index, NewError(errors.New(errors.MissingReturnType), file.path, tokens[:index+1], function)
				}
			}

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
					return function, index, NewError(errors.New(&errors.MissingType{Of: parameterName.Text()}), file.path, tokens[:function.parameterStart+1], function)
				}

				function.Parameters = append(function.Parameters, &Parameter{
					Name:       parameterName.Text(),
					TypeTokens: parameter[1:],
					Position:   function.parameterStart,
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
				return function, index, NewError(errors.New(&errors.MissingType{Of: parameterName.Text()}), file.path, tokens[:function.parameterStart+1], function)
			}

			function.Parameters = append(function.Parameters, &Parameter{
				Name:       parameterName.Text(),
				TypeTokens: parameter[1:],
				Position:   function.parameterStart,
			})

			function.parameterStart = index + 1

		case token.Operator:
			if groupLevel != 0 || function == nil || function.TokenStart != 0 {
				continue
			}

			if t.Text() == "->" {
				function.returnTypeStart = index + 1
			}

		case token.NewLine:
			newlines++

			if newlines == 3 {
				return function, index, NewError(errors.New(errors.UnnecessaryNewlines), file.path, tokens[:index+1], function)
			}
		}
	}

	return function, index, nil
}
