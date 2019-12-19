package errors

// simple is the base class for all errors.
type simple struct {
	Message         string
	RightSideCursor bool
}

func (err *simple) Error() string {
	return err.Message
}

func (err *simple) CursorRight() bool {
	return err.RightSideCursor
}
