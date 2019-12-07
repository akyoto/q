package errors

import "fmt"

// ImportNameAlreadyExists error appears when 2 imported packages have the same base name.
type ImportNameAlreadyExists struct {
	ImportPath string
	Name       string
}

func (err *ImportNameAlreadyExists) Error() string {
	return fmt.Sprintf("Package '%s' has already been imported from '%s'", err.Name, err.ImportPath)
}
