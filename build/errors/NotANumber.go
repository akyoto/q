package errors

import "fmt"

// NotANumber represents number conversion errors.
type NotANumber struct {
	Expression string
}

func (err *NotANumber) Error() string {
	return fmt.Sprintf("Not a number: %s", err.Expression)
}
