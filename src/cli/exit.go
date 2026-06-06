package cli

import (
	"errors"
	"os/exec"

	"git.urbach.dev/cli/q/src/compiler"
)

// Exit codes.
const (
	success = iota
	fail
	invalidArgs
)

// exit returns the exit code depending on the error type.
func exit(err error) int {
	if err == nil {
		return success
	}

	var (
		exit              *exec.ExitError
		expectedParameter *ExpectedParameter
		unknownParameter  *UnknownParameter
		invalidValue      *InvalidValue
		multiError        *compiler.MultiError
	)

	if errors.As(err, &multiError) {
		for _, err := range multiError.Errors {
			showError(err)
		}

		return fail
	}

	showError(err)

	if errors.As(err, &exit) {
		return exit.ExitCode()
	}

	if errors.As(err, &expectedParameter) || errors.As(err, &unknownParameter) || errors.As(err, &invalidValue) {
		return invalidArgs
	}

	return fail
}