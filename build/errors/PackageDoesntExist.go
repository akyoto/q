package errors

import "fmt"

// PackageDoesntExist error appears the imported path doesn't exist on the disk.
type PackageDoesntExist struct {
	ImportPath string
	FilePath   string
}

func (err *PackageDoesntExist) Error() string {
	if err.FilePath == "" {
		return fmt.Sprintf("Package '%s' doesn't exist", err.ImportPath)
	}

	return fmt.Sprintf("Package '%s' doesn't exist in '%s'", err.ImportPath, err.FilePath)
}
