package build

import (
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// IfStart handles the start of if conditions.
func (state *State) IfStart(tokens []token.Token) error {
	state.Expect(token.Keyword)
	expression := tokens[1:]
	variableName := expression[0].Text()
	variable := state.scopes.Get(variableName)

	if variable == nil {
		return &errors.UnknownVariable{VariableName: variableName}
	}

	variable.AliveUntil = state.instrCursor + 1
	temporary, err := state.CompareExpression(variable.Register(), expression[2:], "")

	if err != nil {
		return err
	}

	if temporary != nil {
		temporary.Free()
	}

	endIf := "if_1_end"
	operator := expression[1].Text()

	switch operator {
	case ">=":
		state.assembler.JumpIfLess(endIf)

	case ">":
		state.assembler.JumpIfLessOrEqual(endIf)

	case "<=":
		state.assembler.JumpIfGreater(endIf)

	case "<":
		state.assembler.JumpIfGreaterOrEqual(endIf)

	case "==":
		state.assembler.JumpIfNotEqual(endIf)

	case "!=":
		state.assembler.JumpIfEqual(endIf)
	}

	return nil
}

// IfEnd handles the end of if conditions.
func (state *State) IfEnd() error {
	label := "if_1_end"
	state.assembler.AddLabel(label)
	return nil
}
