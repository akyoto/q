package scanner

import (
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// scanSignature scans only the function signature without the body.
func scanSignature(file *fs.File, pkg string, tokens token.List, i int, delimiter token.Kind) (*core.Function, int, error) {
	var (
		groupLevel  = 0
		nameStart   = i
		inputStart  = -1
		inputEnd    = -1
		outputStart = -1
		outputEnd   = -1
	)

	i++

	// Function parameters
	for i < len(tokens) {
		if tokens[i].Kind == token.GroupStart {
			groupLevel++
			i++

			if groupLevel == 1 {
				inputStart = i
			}

			continue
		}

		if tokens[i].Kind == token.GroupEnd {
			groupLevel--

			if groupLevel < 0 {
				return nil, i, errors.New(MissingGroupStart, file, tokens[i].Position)
			}

			if groupLevel == 0 {
				inputEnd = i
				i++
				break
			}

			i++
			continue
		}

		if tokens[i].Kind == token.Invalid {
			return nil, i, errors.New(&InvalidCharacter{Character: tokens[i].String(file.Bytes)}, file, tokens[i].Position)
		}

		if tokens[i].Kind == token.EOF {
			if groupLevel > 0 {
				return nil, i, errors.New(MissingGroupEnd, file, tokens[i].Position)
			}

			return nil, i, errors.New(InvalidFunctionDefinition, file, tokens[i].Position)
		}

		if groupLevel > 0 {
			i++
			continue
		}

		return nil, i, errors.New(InvalidFunctionDefinition, file, tokens[i].Position)
	}

	// Return type
	if i < len(tokens) && tokens[i].Kind == token.ReturnType {
		outputStart = i + 1

		for i < len(tokens) && tokens[i].Kind != delimiter {
			i++
		}

		outputEnd = i
	}

	name := tokens[nameStart].String(file.Bytes)
	function := core.NewFunction(name, pkg, file)
	parameters := tokens[inputStart:inputEnd]

	for position, param := range parameters.Split {
		if len(param) == 0 {
			return nil, i, errors.New(MissingParameter, file, position)
		}

		if len(param) == 1 {
			return nil, i, errors.New(MissingParameterType, file, param[0].End())
		}

		if param[0].Kind != token.Identifier {
			return nil, i, errors.New(InvalidParameterName, file, param[0].Position)
		}

		function.Input = append(function.Input, &ssa.Parameter{
			Name:   param[0].String(file.Bytes),
			Tokens: param,
			Source: token.Source{StartPos: param[0].Position, EndPos: param[0].End()},
		})
	}

	if outputStart == -1 {
		return function, i, nil
	}

	if tokens[outputStart].Kind == token.GroupStart {
		if tokens[outputEnd-1].Kind == token.GroupEnd {
			outputStart++
			outputEnd--
		} else {
			return nil, i, errors.New(MissingGroupEnd, file, tokens[outputEnd-1].Position)
		}
	}

	outputTokens := tokens[outputStart:outputEnd]

	if len(outputTokens) == 0 {
		return nil, i, errors.New(MissingParameter, file, tokens[outputStart].Position)
	}

	for position, param := range outputTokens.Split {
		if len(param) == 0 {
			return nil, i, errors.New(MissingParameter, file, position)
		}

		if len(param) == 1 {
			function.Output = append(function.Output, &ssa.Parameter{
				Name:   "",
				Tokens: param,
			})
		} else {
			function.Output = append(function.Output, &ssa.Parameter{
				Name:   param[0].String(file.Bytes),
				Tokens: param,
			})
		}
	}

	return function, i, nil
}