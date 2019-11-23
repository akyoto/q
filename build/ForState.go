package build

// ForState handles the state of for loop compilation.
type ForState struct {
	counter   int
	labels    []string
	variables []*Variable
}
