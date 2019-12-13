package errors

import "fmt"

// UnknownField represents unknown variables.
type UnknownField struct {
	TypeName    string
	Name        string
	CorrectName string
}

func (err *UnknownField) Error() string {
	if err.CorrectName != "" {
		return fmt.Sprintf("Type '%s' doesn't have the field '%s', did you mean '%s'?", err.TypeName, err.Name, err.CorrectName)
	}

	return fmt.Sprintf("Type '%s' doesn't have the field '%s'", err.TypeName, err.Name)
}
