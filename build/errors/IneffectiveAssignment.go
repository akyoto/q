package errors

import "fmt"

// IneffectiveAssignment error appears when the value of an assignment is never used.
type IneffectiveAssignment struct {
	Name string
}

func (err *IneffectiveAssignment) Error() string {
	return fmt.Sprintf("This value of '%s' has never been used", err.Name)
}
