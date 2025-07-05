package scanner

import (
	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/x86"
)

// scanFunction scans a function.
func (s *scanner) scanFunction(file *fs.File, tokens token.List, i int) (int, error) {
	function, i, err := scanSignature(file, file.Package, tokens, i, token.BlockStart)

	if err != nil {
		return i, err
	}

	var (
		blockLevel = 0
		bodyStart  = -1
	)

	// Function definition
	for i < len(tokens) {
		if tokens[i].Kind == token.BlockStart {
			blockLevel++
			i++

			if blockLevel == 1 {
				bodyStart = i
			}

			continue
		}

		if tokens[i].Kind == token.BlockEnd {
			blockLevel--

			if blockLevel < 0 {
				return i, errors.New(MissingBlockStart, file, tokens[i].Position)
			}

			if blockLevel == 0 {
				break
			}

			i++
			continue
		}

		if tokens[i].Kind == token.Invalid {
			return i, errors.New(&InvalidCharacter{Character: tokens[i].String(file.Bytes)}, file, tokens[i].Position)
		}

		if tokens[i].Kind == token.EOF {
			if blockLevel > 0 {
				return i, errors.New(MissingBlockEnd, file, tokens[i].Position)
			}

			return i, errors.New(ExpectedFunctionDefinition, file, tokens[i].Position)
		}

		if blockLevel > 0 {
			i++
			continue
		}

		return i, errors.New(ExpectedFunctionDefinition, file, tokens[i].Position)
	}

	switch s.build.Arch {
	case build.ARM:
		switch s.build.OS {
		case build.Linux:
			function.CPU = &arm.LinuxCPU
		case build.Mac:
			function.CPU = &arm.MacCPU
		case build.Windows:
			function.CPU = &arm.WindowsCPU
		}
	case build.X86:
		switch s.build.OS {
		case build.Linux:
			function.CPU = &x86.LinuxCPU
		case build.Mac:
			function.CPU = &x86.MacCPU
		case build.Windows:
			function.CPU = &x86.WindowsCPU
		}
	}

	function.Body = tokens[bodyStart:i]
	s.functions <- function
	return i, nil
}