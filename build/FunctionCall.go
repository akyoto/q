package build

import "github.com/akyoto/q/build/token"

// FunctionCall represents a function call in the source code.
type FunctionCall struct {
	Function   *Function
	Parameters [][]token.Token
}
