package verbose

import (
	"fmt"
	"slices"
	"strings"

	"git.urbach.dev/cli/q/src/fs"
)

// Files shows all imported files.
func Files(files []*fs.File) {
	slices.SortStableFunc(files, func(a *fs.File, b *fs.File) int {
		if a.Package == "main" && b.Package != "main" {
			return -1
		}

		if a.Package != "main" && b.Package == "main" {
			return 1
		}

		return strings.Compare(a.Path, b.Path)
	})

	for _, file := range files {
		fmt.Println(file.Path)
	}
}