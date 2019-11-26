package errors

import "fmt"

// UnknownExpression represents tokenizer failures.
type UnknownExpression struct {
	Expression string
}

func (err *UnknownExpression) Error() string {
	return fmt.Sprintf("Unknown expression: '%s'", err.Expression)
}
