package build

import "github.com/akyoto/q/build/register"

// ForState handles the state of for loop compilation.
type ForState struct {
	counter     int
	labels      []string
	registers   []*register.Register
	temporaries []*register.Register
}
