package errors

// static is used for static errors that have no parameters.
type static struct {
	Message string
}

// Error implements the error interface.
func (err *static) Error() string {
	return err.Message
}