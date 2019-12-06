package build

import (
	"fmt"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/token"
)

// ForState handles the state of for loop compilation.
type ForState struct {
	counter int
	stack   []ForLoop
}

// ForLoop represents a for loop.
type ForLoop struct {
	labelStart string
	labelEnd   string
	counter    *register.Register
	limit      *register.Register
}

// ForStart handles the start of for loops.
func (state *State) ForStart(tokens []token.Token) error {
	state.Expect(token.Keyword)
	state.scopes.Push()
	expression := tokens[1:]

	rangePos := token.Index(expression, token.Range)

	if rangePos == -1 {
		return errors.MissingRange
	}

	operatorPos := token.Index(expression, token.Operator)
	var register *register.Register

	if operatorPos == -1 {
		start := expression[:rangePos]

		if len(start) == 0 {
			return errors.MissingRangeStart
		}

		register = state.registers.General.FindFree()

		if register == nil {
			return errors.ExceededMaxVariables
		}

		err := state.TokensToRegister(start, register)

		if err != nil {
			return err
		}
	} else {
		assignment := expression[:rangePos]
		variable, err := state.AssignVariable(assignment)

		if err != nil {
			return err
		}

		register = variable.Register()
	}

	state.forState.counter++

	labelStart := fmt.Sprintf("for_%d", state.forState.counter)
	labelEnd := fmt.Sprintf("for_%d_end", state.forState.counter)

	upperLimit := expression[rangePos+1:]

	if len(upperLimit) == 0 {
		return errors.MissingRangeLimit
	}

	state.tokenCursor++
	temporary, err := state.CompareRegisterExpression(register, upperLimit, labelStart)

	if err != nil {
		return err
	}

	state.assembler.JumpIfEqual(labelEnd)

	state.forState.stack = append(state.forState.stack, ForLoop{
		labelStart: labelStart,
		labelEnd:   labelEnd,
		counter:    register,
		limit:      temporary,
	})

	return nil
}

// ForEnd handles the end of for loops.
func (state *State) ForEnd() error {
	err := state.PopScope(true)

	if err != nil {
		return err
	}

	loop := state.forState.stack[len(state.forState.stack)-1]
	state.forState.stack = state.forState.stack[:len(state.forState.stack)-1]

	state.assembler.IncreaseRegister(loop.counter)
	state.assembler.Jump(loop.labelStart)
	state.assembler.AddLabel(loop.labelEnd)
	loop.counter.Free()

	if loop.limit != nil {
		loop.limit.Free()
	}

	return nil
}
