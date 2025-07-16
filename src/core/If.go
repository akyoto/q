package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// If compiles a branch instruction.
func (f *Function) If(tokens token.List) error {
	blockStart, blockEnd, err := f.block(tokens)

	if err != nil {
		return err
	}

	condExpr := expression.Parse(tokens[1:blockStart])
	cond, err := f.eval(condExpr)

	if err != nil {
		return err
	}

	branch := &ssa.If{
		Condition: cond,
	}

	f.Append(branch)
	trueBlock := f.AddBlock("true")
	branch.Then = trueBlock
	f.compileTokens(tokens[blockStart+1 : blockEnd])

	falseBlock := f.AddBlock("false")
	branch.Else = falseBlock
	return nil
}