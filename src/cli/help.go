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
		case strings.HasPrefix(line, "    "):
			ansi.Green.Print(line)
		case line == "EOF":
		default:
			fmt.Print(line)
		}
	}

	return 0
}