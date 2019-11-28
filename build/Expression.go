package build

import (
	"fmt"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/expression"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/token"
)

// TokenToRegister moves a token into a register.
// It only works with identifiers, numbers and texts.
func (state *State) TokenToRegister(singleToken token.Token, register *register.Register) error {
	switch singleToken.Kind {
	case token.Identifier:
		variableName := singleToken.Text()
		variable := state.scopes.Get(variableName)

		if variable == nil {
			return fmt.Errorf("Unknown variable %s", variableName)
		}

		variable.AliveUntil = state.instrCursor + 1

		// Moving a variable into its own register is pointless
		if variable.Register == register {
			return nil
		}

		state.assembler.MoveRegisterRegister(register, variable.Register)

	case token.Number:
		numberString := singleToken.Text()
		number, err := state.ParseInt(numberString)

		if err != nil {
			return err
		}

		state.assembler.MoveRegisterNumber(register, uint64(number))

	case token.Text:
		address := state.assembler.AddString(singleToken.Text())
		state.assembler.MoveRegisterAddress(register, address)
	}

	if register.IsFree() {
		register.Use(singleToken)
	}

	return nil
}

// TokensToRegister moves the result of a token expression into the given register.
func (state *State) TokensToRegister(tokens []token.Token, register *register.Register) error {
	if len(tokens) == 1 {
		return state.TokenToRegister(tokens[0], register)
	}

	expr, err := expression.FromTokens(tokens)

	if err != nil {
		return err
	}

	err = state.ExpressionToRegister(expr, register)
	expr.Close()
	return err
}

// ExpressionToRegister moves the result of an expression into the given register.
func (state *State) ExpressionToRegister(expr *expression.Expression, finalRegister *register.Register) error {
	if expr.IsLeaf() {
		return state.TokenToRegister(expr.Token, finalRegister)
	}

	expr.SortByRegisterCount()
	expr.Register = finalRegister

	// Assign final register to each left operand
	_ = expr.EachOperation(func(sub *expression.Expression) error {
		left := sub.Children[0]
		left.Register = finalRegister
		return nil
	})

	err := expr.EachOperation(func(sub *expression.Expression) error {
		if sub.IsFunctionCall {
			functionName := sub.Token.Text()
			function := state.environment.Functions[functionName]

			if function == nil {
				return state.UnknownFunctionError(functionName)
			}

			function.Used = true

			// Move parameters into registers
			for i, parameter := range sub.Children {
				parameter.Register = state.registers.Call[i]
				err := state.ExpressionToRegister(parameter, parameter.Register)

				if err != nil {
					return err
				}
			}

			// Call the function
			state.assembler.Call(functionName)

			for _, callRegister := range state.registers.Call {
				callRegister.Free()
			}

			returnValueRegister := state.registers.ReturnValue[0]
			returnValueRegister.Use(sub)

			if sub.Register == nil {
				sub.Register = state.registers.FindFreeRegister()
			}

			// Save return value in temporary register
			if sub.Register != returnValueRegister {
				state.assembler.MoveRegisterRegister(sub.Register, returnValueRegister)
			}

			returnValueRegister.Free()
			return nil
		}

		left := sub.Children[0]
		right := sub.Children[1]

		if left.Register == nil {
			left.Register = state.registers.FindFreeRegister()
		}

		if sub.Token.Kind != token.Separator {
			sub.Register = left.Register

			if sub.Register.IsFree() {
				sub.Register.Use(sub)
			}
		}

		// Left operand
		if left.IsLeaf() {
			err := state.TokenToRegister(left.Token, sub.Register)

			if err != nil {
				return err
			}
		} else if sub.Register != left.Register {
			state.assembler.MoveRegisterRegister(sub.Register, left.Register)
			left.Register.Free()
		}

		operator := sub.Token.Text()

		if operator == "," {
			if right.IsLeaf() {
				err := state.TokenToRegister(right.Token, right.Register)
				return err
			}

			return nil
		}

		// Right operand is a leaf node
		if right.IsLeaf() {
			switch right.Token.Kind {
			case token.Identifier:
				variableName := right.Token.Text()
				variable := state.scopes.Get(variableName)

				if variable == nil {
					return fmt.Errorf("Unknown variable %s", variableName)
				}

				variable.AliveUntil = state.instrCursor + 1
				return state.CalculateRegisterRegister(operator, sub.Register, variable.Register)

			case token.Number:
				return state.CalculateRegisterNumber(operator, sub.Register, right.Token.Text())

			default:
				return fmt.Errorf("Invalid operand %s", right.Token)
			}
		}

		// Right operand is an expression
		err := state.CalculateRegisterRegister(operator, sub.Register, right.Register)

		if right.Register != nil {
			right.Register.Free()
		}

		return err
	})

	if err != nil {
		return err
	}

	// Free registers
	_ = expr.EachOperation(func(expr *expression.Expression) error {
		if expr.Register != nil && expr.Register != finalRegister {
			expr.Register.Free()
		}

		return nil
	})

	// Mark final register as used if it's not marked already
	if finalRegister.IsFree() {
		finalRegister.Use(expr)
	}

	return nil
}

// CalculateRegisterNumber performs an operation on a register and a number.
func (state *State) CalculateRegisterNumber(operation string, register *register.Register, operand string) error {
	number, err := state.ParseInt(operand)

	if err != nil {
		return err
	}

	switch operation {
	case "+":
		if number == 1 && state.optimize {
			state.assembler.IncreaseRegister(register)
			return nil
		}

		state.assembler.AddRegisterNumber(register, uint64(number))

	case "-":
		if number == 1 && state.optimize {
			state.assembler.DecreaseRegister(register)
			return nil
		}

		state.assembler.SubRegisterNumber(register, uint64(number))

	case "*":
		state.assembler.MulRegisterNumber(register, uint64(number))

	case ",":
		return nil

	default:
		return errors.NotImplemented
	}

	return nil
}

// CalculateRegisterRegister performs an operation on two registers.
func (state *State) CalculateRegisterRegister(operation string, registerTo *register.Register, registerFrom *register.Register) error {
	switch operation {
	case "+":
		state.assembler.AddRegisterRegister(registerTo, registerFrom)

	case "-":
		state.assembler.SubRegisterRegister(registerTo, registerFrom)

	case "*":
		state.assembler.MulRegisterRegister(registerTo, registerFrom)

	case ",":
		return nil

	default:
		return errors.NotImplemented
	}

	return nil
}
