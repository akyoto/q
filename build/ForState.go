package build

import "github.com/akyoto/q/build/register"

// ForState handles the state of for loop compilation.
type ForState struct {
	counter int
	stack   []ForLoop
}

// ForLoop represents a for loop.
type ForLoop struct {
	labelStart string
	labelEnd   string
	counter    *register.Register
	limit      *register.Register
}
