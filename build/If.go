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

	label := "if_1_end"
	state.assembler.CompareRegisterNumber(variable.Register(), uint64(number))
	operator := expression[1].Text()

	switch operator {
	case ">=":
		state.assembler.JumpIfLess(label)

	case ">":
		state.assembler.JumpIfLessOrEqual(label)

	case "<=":
		state.assembler.JumpIfGreater(label)

	case "<":
		state.assembler.JumpIfGreaterOrEqual(label)

	case "==":
		state.assembler.JumpIfNotEqual(label)

	case "!=":
		state.assembler.JumpIfEqual(label)
	}

	return nil
}

// IfEnd handles the end of if conditions.
func (state *State) IfEnd() error {
	label := "if_1_end"
	state.assembler.AddLabel(label)
	return nil
}
