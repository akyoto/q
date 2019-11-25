package errors

// Base is the base class for all errors.
type Base struct {
	Message         string
	RightSideCursor bool
}

func (err *Base) Error() string {
	return err.Message
}

func (err *Base) CursorRight() bool {
	return err.RightSideCursor
}
