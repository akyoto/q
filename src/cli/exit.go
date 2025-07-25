package cli

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	fe "git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/go/color"
	"git.urbach.dev/go/color/ansi"
)

// exit returns the exit code depending on the error type.
func exit(err error) int {
	if err == nil {
		return 0
	}

	var (
		exit              *exec.ExitError
		expectedParameter *ExpectedParameter
		unknownParameter  *UnknownParameter
		invalidValue      *InvalidValue
		fileError         *fe.FileError
	)

	if errors.As(err, &fileError) {
		line, offset := fileError.Line()
		indent := strings.Repeat(" ", offset)
		color.Redirect(os.Stderr)
		ansi.Reset.Printf("%s\n\n", fileError.Location())
		ansi.Reset.Printf("    %s\n", line)
		ansi.Red.Printf("%s    ┬\n%s    ╰─ ", indent, indent)
		ansi.Reset.Printf("%s\n\n", fileError.Error())
		ansi.Dim.Println(fileError.Stack())
	} else {
		fmt.Fprintln(os.Stderr, err)
	}

	if errors.As(err, &exit) {
		return exit.ExitCode()
	}

	if errors.As(err, &expectedParameter) || errors.As(err, &unknownParameter) || errors.As(err, &invalidValue) {
		return 2
	}

	return 1
}