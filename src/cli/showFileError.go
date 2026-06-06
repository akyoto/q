package cli

import (
	"os"
	"strings"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/go/color"
	"git.urbach.dev/go/color/ansi"
)

// showFileError shows an error on stderr with file, line and column information.
func showFileError(fileError *errors.FileError) {
	line, offset := fileError.Line()
	indent := strings.Repeat(" ", offset)
	color.Redirect(os.Stderr)
	ansi.Reset.Printf("%s\n\n", fileError.Link())
	source := fileError.Source()
	length := int(source.End() - source.Start())

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