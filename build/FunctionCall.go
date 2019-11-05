package build

import "github.com/akyoto/q/spec"

// FunctionCall
type FunctionCall struct {
	Function        *spec.Function
	Parameters      []Expression
	ProcessedTokens int
}

// Reset resets the state before sending this object back into the memory pool.
func (call *FunctionCall) Reset() {
	call.Function = nil
	call.Parameters = call.Parameters[:0]
	call.ProcessedTokens = 0
}
