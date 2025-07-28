package verbose

import (
	"fmt"
	"strings"

	"git.urbach.dev/go/color"
)

// Header shows an ASCII art banner with a little too many colors.
func Header(banner string) {
	totalLineCount := color.Value(strings.Count(banner, "\n"))
	lineCount := 0

	for line := range strings.Lines(banner) {
		y := color.Value(lineCount) / totalLineCount
		characters := color.Value(len(line) - 1)

		for i, c := range line {
			x := color.Value(i) / characters
			color.HSL(180+x*180+y*180, 1.0, 0.7).Print(string(c))
		}

		lineCount++
	}

	fmt.Print("\n\n")
}