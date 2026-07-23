package cli

import (
	"fmt"
	"runtime/debug"
	"strings"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/go/color/ansi"
)

// render parses the text and prints it with colors.
func render(text string) {
	for line := range strings.Lines(text) {
		start := strings.IndexByte(line, '{')

		for start != -1 {
			end := strings.IndexByte(line, '}')
			key := line[start+1 : end]
			value := buildSetting(key)
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
			if strings.HasSuffix(line, " options:\n") {
				fmt.Print(line)
			} else {
				ansi.Green.Print(line)
			}
		case line == "EOF":
		default:
			fmt.Print(line)
		}

		if len(dimmed) > 0 {
			ansi.Dim.Print(dimmed)
		}
	}
}

// buildSetting retrieves information about the build.
func buildSetting(key string) string {
	switch key {
	case "arch":
		return config.New().Arch.String()
	case "commit":
		hash := debugBuildSetting("vcs.revision")

		if len(hash) >= 7 {
			return hash[:7]
		}

		return hash
	case "date":
		date := debugBuildSetting("vcs.time")

		if len(date) >= 10 {
			return date[:10]
		}

		return date
	case "os":
		return config.New().OS.String()
	default:
		return ""
	}
}

// debugBuildSetting retrieves build data that Go saved within the binary.
func debugBuildSetting(key string) string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == key {
				return setting.Value
			}
		}
	}

	return ""
}