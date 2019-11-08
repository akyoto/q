package build

import "github.com/akyoto/q/instruction"

// FunctionCall represents a function call in the source code.
type FunctionCall struct {
	Function       *Function
	Parameters     []instruction.Expression
	ParameterStart int
}

// Reset resets the state before sending this object back into the memory pool.
func (call *FunctionCall) Reset() {
	call.Function = nil
	call.Parameters = call.Parameters[:0]
	call.ParameterStart = 0
}
