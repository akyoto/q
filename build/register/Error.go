package register

import "fmt"

// ErrAlreadyInUse type errors are returned when a register is already in use.
type ErrAlreadyInUse struct {
	Register *Register
	UsedBy   fmt.Stringer
}

// Error implements the text representation.
func (err *ErrAlreadyInUse) Error() string {
	return fmt.Sprintf("Register '%s' already used by '%s'", err.Register.Name, err.UsedBy)
}
