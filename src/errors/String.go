package errors

// String is used for static errors that have no parameters.
type String struct {
	Message string
}

// Error implements the error interface.
func (err *String) Error() string {
	return err.Message
}