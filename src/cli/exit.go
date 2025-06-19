package cli

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// exit returns the exit code depending on the error type.
func exit(err error) int {
	fmt.Fprintln(os.Stderr, err)

	var (
		exit              *exec.ExitError
		expectedParameter *expectedParameterError
		unknownParameter  *unknownParameterError
		invalidValue      *invalidValueError
	)

	if errors.As(err, &exit) {
		return exit.ExitCode()
	}

	if errors.As(err, &expectedParameter) || errors.As(err, &unknownParameter) || errors.As(err, &invalidValue) {
		return 2
	}

	return 1
}