package errors

import "fmt"

// CantInferType error appears when a type can't be inferred.
type CantInferType struct {
	Expression string
}

func (err *CantInferType) Error() string {
	return fmt.Sprintf("Can't infer type of expression '%s'", err.Expression)
}
