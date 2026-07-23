package cli

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"git.urbach.dev/cli/q/src/compiler"
	fe "git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/go/color"
	"git.urbach.dev/go/color/ansi"
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

// showError shows an error on stderr.
func showError(err error) {
	var fileError *fe.FileError

	if errors.As(err, &fileError) {
		showFileError(fileError)
	} else {
		fmt.Fprintln(os.Stderr, err)
	}
}

// showFileError shows an error on stderr with file, line and column information.
func showFileError(fileError *fe.FileError) {
	line, offset := fileError.Line()
	indent := strings.Repeat(" ", offset)
	color.Redirect(os.Stderr)
	ansi.Reset.Printf("%s\n\n", fileError.Link())
	source := fileError.Source()
	length := int(source.End() - source.Start())

	if offset+length > len(line) {
		length = len(line) - offset
	}

	if length > 0 {
		ansi.Reset.Printf("    %s", line[:offset])
		ansi.Red.Print(line[offset : offset+length])
		ansi.Reset.Println(line[offset+length:])
	} else {
		ansi.Reset.Printf("    %s\n", line)
	}

	ansi.Red.Printf("%s    ┬\n", indent)
	ansi.Red.Printf("%s    ╰─ ", indent)
	ansi.Reset.Printf("%s\n\n", fileError.Error())
	ansi.Dim.Println(fileError.Stack())
}

// invalid shows the help on stderr and returns exit code 2 (invalid parameters).
func invalid() int {
	color.Redirect(os.Stderr)
	help()
	return invalidArgs
}