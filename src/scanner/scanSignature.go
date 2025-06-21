package scanner

import (
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// scanSignature scans only the function signature without the body.
func scanSignature(file *fs.File, tokens token.List, i int, delimiter token.Kind) (*core.Function, int, error) {
	var (
		groupLevel  = 0
		nameStart   = i
		paramsStart = -1
		paramsEnd   = -1
		typeStart   = -1
		typeEnd     = -1
	)

	i++

	// Function parameters
	for i < len(tokens) {
		if tokens[i].Kind == token.GroupStart {
			groupLevel++
			i++

			if groupLevel == 1 {
				paramsStart = i
			}

			continue
		}

		if tokens[i].Kind == token.GroupEnd {
			groupLevel--

			if groupLevel < 0 {
				return nil, i, errors.New(MissingGroupStart, file, tokens[i].Position)
			}

			if groupLevel == 0 {
				paramsEnd = i
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

			if paramsStart == -1 {
				return nil, i, errors.New(InvalidFunctionDefinition, file, tokens[i].Position)
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
		typeStart = i + 1

		for i < len(tokens) && tokens[i].Kind != delimiter {
			i++
		}

		typeEnd = i
	}

	name := tokens[nameStart].String(file.Bytes)
	function := core.NewFunction(name, file)
	parameters := tokens[paramsStart:paramsEnd]

	for param := range parameters.Split {
		if len(param) == 0 {
			return nil, i, errors.New(MissingParameter, file, parameters[0].Position)
		}

		if len(param) == 1 {
			return nil, i, errors.New(MissingType, file, param[0].End())
		}

		function.Input = append(function.Input, &core.Parameter{
			Name:       param[0].String(file.Bytes),
			TypeTokens: param[1:],
		})
	}

	if typeStart != -1 {
		if tokens[typeStart].Kind == token.GroupStart && tokens[typeEnd-1].Kind == token.GroupEnd {
			typeStart++
			typeEnd--
		}

		outputTokens := tokens[typeStart:typeEnd]

		if len(outputTokens) == 0 {
			return nil, i, errors.New(MissingParameter, file, tokens[typeStart].Position)
		}

		errorPos := token.Position(typeStart)

		for param := range outputTokens.Split {
			if len(param) == 0 {
				return nil, i, errors.New(MissingParameter, file, errorPos)
			}

			if len(param) == 1 {
				function.Output = append(function.Output, &core.Parameter{
					Name:       "",
					TypeTokens: param,
				})
			} else {
				function.Output = append(function.Output, &core.Parameter{
					Name:       param[0].String(file.Bytes),
					TypeTokens: param[1:],
				})
			}

			errorPos = param[len(param)-1].End() + 1
		}
	}

	return function, i, nil
}