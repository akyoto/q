package cli

import (
	"fmt"
	"runtime/debug"
	"strings"

	"git.urbach.dev/go/color/ansi"
)

// show parses the text and prints it with colors.
func show(text string) {
	for line := range strings.Lines(text) {
		start := strings.IndexByte(line, '{')

		for start != -1 {
			end := strings.IndexByte(line, '}')
			key := line[start+1 : end]
			value := ""

			switch key {
			case "commit":
				value = buildInfo("vcs.revision")
			case "date":
				value = buildInfo("vcs.time")
			}

			line = line[:start] + value + line[end+1:]
			start = strings.IndexByte(line, '{')
		}

		startDimmed := strings.IndexByte(line, '#')
		dimmed := ""

		if startDimmed != -1 {
			dimmed = line[startDimmed:]
			line = line[:startDimmed]
		}

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

		if len(dimmed) > 0 {
			ansi.Dim.Print(dimmed)
		}
	}
}

// buildInfo retrieves information about the build.
func buildInfo(key string) string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == key {
				return setting.Value
			}
		}
	}

	return ""
}