package core

import (
	"git.urbach.dev/cli/q/src/token"
)

// compileTokens compiles a token list.
func (f *Function) compileTokens(tokens token.List) error {
	for instr := range tokens.Instructions {
		err := f.compileInstruction(instr)

		if err != nil {
			return err
		}
	}

	return nil
}