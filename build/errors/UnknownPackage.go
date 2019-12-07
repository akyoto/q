package errors

import "fmt"

// UnknownPackage represents unknown packages.
type UnknownPackage struct {
	Name        string
	CorrectName string
}

func (err *UnknownPackage) Error() string {
	if err.CorrectName != "" {
		return fmt.Sprintf("Unknown package '%s', did you mean '%s'?", err.Name, err.CorrectName)
	}

	return fmt.Sprintf("Unknown package '%s'", err.Name)
}
