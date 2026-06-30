package cli

import (
	"fmt"
	"slices"
	"strings"

	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/color/ansi"
)

// format formats source files.
func format(args []string) int {
	noSpaceLeft := []token.Kind{
		token.Dot,
		token.ArrayStart,
		token.ArrayEnd,
		token.GroupStart,
		token.GroupEnd,
		token.Separator,
		token.FieldAssign,
		token.Range,
	}

	noSpaceRight := []token.Kind{
		token.Dot,
		token.ArrayStart,
		token.ArrayEnd,
		token.GroupStart,
		token.GroupEnd,
		token.NewLine,
		token.Range,
	}

	indent := 0
	indented := false
	lastKind := token.NewLine

	for _, arg := range args {
		contents, err := fs.ReadFile(arg)

		if err != nil {
			return exit(err)
		}

		tokens := token.Tokenize(contents)

		for i, t := range tokens {
			if t.Kind == token.NewLine {
				fmt.Println()
				indented = false
				lastKind = t.Kind
				continue
			}

			if t.Kind == token.BlockStart {
				indent++
			}

			if t.Kind == token.BlockEnd {
				indent--
			}

			if !indented {
				fmt.Print(strings.Repeat("\t", indent))
				indented = true
			}

			c := ansi.Reset

			switch {
			case t.Kind == token.Identifier && lastKind == token.NewLine && indent == 0:
				c = ansi.Yellow
			case t.Kind == token.Number:
				c = ansi.Cyan
			case t.Kind == token.String || t.Kind == token.Rune:
				c = ansi.Green
			case t.Kind == token.Comment:
				c = ansi.Dim
			case t.Kind.IsKeyword():
				c = ansi.Red
			}

			switch {
			case t.Kind == token.Define || lastKind == token.Define:
				fmt.Print(" ")
			case t.Kind == token.ReturnType || lastKind == token.ReturnType:
				fmt.Print(" ")
			case t.Kind == token.BlockStart && lastKind == token.GroupEnd:
				fmt.Print(" ")
			case t.Kind == token.BlockEnd && lastKind == token.GroupEnd:
				fmt.Print(" ")
			case t.Kind == token.BlockStart && lastKind == token.Identifier && i+1 < len(tokens) && tokens[i+1].Kind != token.NewLine:
				// no space
			case slices.Contains(noSpaceLeft, t.Kind) || slices.Contains(noSpaceRight, lastKind):
				// no space
			case t.Kind.IsAssignment() || lastKind.IsAssignment():
				fmt.Print(" ")
			default:
				fmt.Print(" ")
			}

			lastKind = t.Kind
			c.Print(t.StringFrom(contents))
		}
	}

	return success
}