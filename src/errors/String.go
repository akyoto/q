package errors

// String creates a static error message without parameters.
func String(message string) *static {
	return &static{Message: message}
}