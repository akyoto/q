package cli

import (
	"errors"
	"fmt"
	"os"

	fe "git.urbach.dev/cli/q/src/errors"
)

// showError shows an error on stderr.
func showError(err error) {
	var fileError *fe.FileError

	if errors.As(err, &fileError) {
		showFileError(fileError)
	} else {
		fmt.Fprintln(os.Stderr, err)
	}
}