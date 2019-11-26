package build

import "github.com/akyoto/q/build/token"

// Call represents a function call in the source code.
type Call struct {
	Function   *Function
	Parameters [][]token.Token
}
