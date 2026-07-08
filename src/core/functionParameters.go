package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// functionParameters contains the input and output parameters of a function.
type functionParameters struct {
	Input  []*ssa.Parameter
	Output []*ssa.Parameter
}

// AddInput adds an input parameter.
func (p *functionParameters) AddInput(tokens token.List, source token.Source) {
	p.Input = append(p.Input, &ssa.Parameter{
		Tokens: tokens,
		Source: source,
	})
}

// AddOutput adds an output parameter.
func (p *functionParameters) AddOutput(tokens token.List, source token.Source) {
	p.Output = append(p.Output, &ssa.Parameter{
		Tokens: tokens,
		Source: source,
	})
}