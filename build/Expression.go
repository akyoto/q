package build

import (
	"fmt"
	"sync/atomic"

	"github.com/akyoto/asm/syscall"
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/expression"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/token"
)

// ExpressionToRegister moves the result of an expression into the given register.
func (state *State) ExpressionToRegister(root *expression.Expression, finalRegister *register.Register) error {
	if root.IsLeaf() {
		return state.TokenToRegister(root.Token, finalRegister)
	}

	// Save the temporary registers so we can easily free them later
	var temporaryRegisters []*register.Register

	// Sort by expression complexity so that we can
	// calculate the most complex expression first.
	// This reduces the number of registers required.
	root.SortByRegisterCount()
	root.Register = finalRegister

	// Assign final register to the left operands in the left tree
	left := root

	for len(left.Children) > 0 {
		left = left.Children[0]
		left.Register = finalRegister
	}

	// Execute each operation starting from the bottom left
	err := root.EachOperation(func(sub *expression.Expression) error {
		if sub.IsFunctionCall {
			// Allocate a temporary register if necessary
			if sub.Register == nil {
				sub.Register = state.registers.FindFreeRegister()
				sub.Register.Use(sub)
				temporaryRegisters = append(temporaryRegisters, sub.Register)
			}

			return state.CallExpression(sub)
		}

		left := sub.Children[0]
		right := sub.Children[1]

		// Allocate a temporary register if necessary
		if left.Register == nil {
			left.Register = state.registers.FindFreeRegister()
			left.Register.Use(sub)
			temporaryRegisters = append(temporaryRegisters, left.Register)
		}

		sub.Register = left.Register

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
				return state.CalculateRegisterRegister(operator, sub.Register, variable.Register())

			case token.Number:
				return state.CalculateRegisterNumber(operator, sub.Register, right.Token.Text())

			default:
				return fmt.Errorf("Invalid operand %s", right.Token)
			}
		}

		// Right operand is an expression
		return state.CalculateRegisterRegister(operator, sub.Register, right.Register)
	})

	if err != nil {
		return err
	}

	// Free temporary registers
	for _, reg := range temporaryRegisters {
		reg.Free()
	}

	// Mark final register as used if it's not marked already
	if finalRegister != nil && finalRegister.IsFree() {
		finalRegister.Use(root)
	}

	return nil
}

// CallExpression executes a function call.
func (state *State) CallExpression(expr *expression.Expression) error {
	functionName := expr.Token.Text()
	function := state.environment.Functions[functionName]
	isBuiltin := false

	if function == nil {
		function = BuiltinFunctions[functionName]
		isBuiltin = true
	}

	if function == nil {
		return state.UnknownFunctionError(functionName)
	}

	parameters := expr.Children

	// Calling a function with side effects causes our function to have side effects
	if atomic.LoadInt32(&function.SideEffects) > 0 {
		atomic.AddInt32(&state.function.SideEffects, 1)
	}

	// Parameter check
	if !function.NoParameterCheck && len(parameters) != len(function.Parameters) {
		return &errors.ParameterCount{
			FunctionName:  function.Name,
			CountGiven:    len(parameters),
			CountRequired: len(function.Parameters),
		}
	}

	// print is a little special
	if isBuiltin && functionName == "print" {
		parameter := parameters[0]

		if parameter.Token.Kind != token.Text {
			return fmt.Errorf("'%s' requires a text parameter instead of '%s'", function.Name, parameter.Token.Text())
		}

		text := parameter.Token.Text() + "\n"
		address := state.assembler.AddString(text)
		state.assembler.MoveRegisterNumber(state.registers.Syscall[0], uint64(syscall.Write))
		state.assembler.MoveRegisterNumber(state.registers.Syscall[1], 1)
		state.assembler.MoveRegisterAddress(state.registers.Syscall[2], address)
		state.assembler.MoveRegisterNumber(state.registers.Syscall[3], uint64(len(text)))
		state.assembler.Syscall()
		return nil
	}

	// Call the function
	err := state.BeforeCall(parameters)

	if err != nil {
		return err
	}

	if functionName == "syscall" {
		state.assembler.Syscall()
	} else {
		state.assembler.Call(functionName)
	}

	state.AfterCall(function)

	// Free the call registers
	for _, callRegister := range state.registers.Call {
		callRegister.Free()
	}

	// Mark return value register temporarily as used for better assembly output
	returnValueRegister := state.registers.ReturnValue[0]
	returnValueRegister.Use(expr)

	// Save return value in temporary register
	if expr.Register != returnValueRegister {
		state.assembler.MoveRegisterRegister(expr.Register, returnValueRegister)
		returnValueRegister.Free()
	}

	return nil
}

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
		if variable.Register() == register {
			return nil
		}

		state.assembler.MoveRegisterRegister(register, variable.Register())

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
