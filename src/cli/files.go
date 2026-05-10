package cli

import (
	"fmt"
	"slices"
	"strings"

	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/scanner"
)

// files shows the entire list of files that are used in a build.
func files(args []string) int {
	b, err := newBuild(args)

	if err != nil {
		return exit(err)
	}

	env, err := scanner.Scan(b)

	if err != nil {
		return exit(err)
	}

	slices.SortStableFunc(env.Files, func(a *fs.File, b *fs.File) int {
		if a.Package == "main" && b.Package != "main" {
			return -1
		}

		if a.Package != "main" && b.Package == "main" {
			return 1
		}

		return strings.Compare(a.Path, b.Path)
	})

	for _, file := range env.Files {
		fmt.Println(file.Path)
	}

	return exit(nil)
}