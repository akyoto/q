package errors

import (
	"fmt"
)

// VariableAlreadyExists is used when existing variables are used for new variable declarations.
type VariableAlreadyExists struct {
	Name string
}

func (err *VariableAlreadyExists) Error() string {
	return fmt.Sprintf("Variable '%s' already exists", err.Name)
}
