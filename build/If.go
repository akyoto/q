package build

import (
	"fmt"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/operators"
	"github.com/akyoto/q/build/token"
)

// IfState handles the state of branch compilation.
type IfState struct {
	counter int
	labels  []string
}

// IfStart handles the start of if conditions.
func (state *State) IfStart(tokens []token.Token) error {
	state.Skip(token.Keyword)
	state.scopes.Push()
	condition := tokens[1:]

	state.ifState.counter++
	elseLabel := fmt.Sprintf("if_%d_end", state.ifState.counter)
	state.ifState.labels = append(state.ifState.labels, elseLabel)

	return state.Condition(condition, elseLabel)
}

// Condition encodes a compare instruction for the given condition.
func (state *State) Condition(condition []token.Token, elseLabel string) error {
	operatorPos := -1

	for i, t := range condition {
		if t.Kind == token.Operator && operators.All[t.Text()].Kind == operators.Comparison {
			operatorPos = i
			break
		}
	}

	if operatorPos == -1 {
		return errors.New(errors.InvalidExpression)
	}

	left := condition[:operatorPos]
	leftRegister, leftType, err := state.EvaluateTokens(left)

	if err != nil {
		return err
	}

	if leftType == nil {
		return errors.New(&errors.CantInferType{Expression: fmt.Sprint(left)})
	}

	right := condition[operatorPos+1:]
	temporary, rightType, err := state.CompareRegisterExpression(leftRegister, right, "")

	if err != nil {
		return err
	}

	if rightType == nil {
		return errors.New(&errors.CantInferType{Expression: fmt.Sprint(right)})
	}

	if leftType != rightType {
		return errors.New(&errors.InvalidType{Type: rightType.Name, Expected: leftType.Name})
	}

	if temporary != nil {
		temporary.Free()
	}

	operator := condition[operatorPos].Text()
	state.IfFalseJump(operator, elseLabel)
	return nil
}

// IfFalseJump jumps if the previous compare statement was false.
func (state *State) IfFalseJump(operator string, label string) {
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
}

// IfEnd handles the end of if conditions.
func (state *State) IfEnd() error {
	err := state.PopScope(false)

	if err != nil {
		return err
	}

	label := state.ifState.labels[len(state.ifState.labels)-1]
	state.ifState.labels = state.ifState.labels[:len(state.ifState.labels)-1]
	state.assembler.AddLabel(label)
	return nil
}
