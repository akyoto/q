package cli

import (
	_ "embed"
	"fmt"
	"strings"

	"git.urbach.dev/go/color/ansi"
)

//go:embed help.txt
var helpText string

// help shows the command line argument usage.
func help() int {
	for line := range strings.Lines(helpText) {
		switch {
		case strings.HasPrefix(line, "        "):
			startDimmed := strings.IndexByte(line, '#')
			dimmed := ""

			if startDimmed != -1 {
				dimmed = line[startDimmed:]
				line = line[:startDimmed]
			}

			for {
				start := strings.IndexByte(line, '[')

				if start == -1 {
					ansi.Cyan.Print(line)
					break
				}

				end := strings.IndexByte(line, ']')
				ansi.Cyan.Print(line[:start])
				ansi.Yellow.Print(line[start : end+1])
				line = line[end+1:]
			}

			if len(dimmed) > 0 {
				ansi.Dim.Print(dimmed)
			}
		case strings.HasPrefix(line, "    "):
			ansi.Green.Print(line)
		case line == "EOF":
		default:
			fmt.Print(line)
		}
	}

	return 0
}