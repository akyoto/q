package build

// FunctionCall
type FunctionCall struct {
	Function        *Function
	Parameters      []Expression
	ProcessedTokens int
}

// Reset resets the state before sending this object back into the memory pool.
func (call *FunctionCall) Reset() {
	call.Function = nil
	call.Parameters = call.Parameters[:0]
	call.ProcessedTokens = 0
}
