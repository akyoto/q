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

	numberString := expression[len(expression)-1].Text()
	number, err := state.ParseInt(numberString)

	if err != nil {
		return err
	}

	endIf := "if_1_end"
	state.assembler.CompareRegisterNumber(variable.Register(), uint64(number))
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
