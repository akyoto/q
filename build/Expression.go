package build

import (
	"fmt"
	"sync/atomic"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/expression"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/token"
	"github.com/akyoto/q/build/types"
	"github.com/akyoto/stringutils/unsafe"
)

// EvaluateTokens evaluates the token expression and stores the result in a register.
func (state *State) EvaluateTokens(tokens []token.Token) (*register.Register, *types.Type, error) {
	freeRegister := state.registers.General.FindFree()

	if freeRegister == nil {
		return nil, nil, errors.New(errors.ExceededMaxVariables)
	}

	typ, err := state.TokensToRegister(tokens, freeRegister)
	return freeRegister, typ, err
}

// TokensToRegister moves the result of a token expression into the given register.
func (state *State) TokensToRegister(tokens []token.Token, register *register.Register) (*types.Type, error) {
	if len(tokens) == 1 {
		return state.TokenToRegister(tokens[0], register)
	}

	expr, err := expression.FromTokens(tokens)

	if err != nil {
		return nil, err
	}

	typ, err := state.ExpressionToRegister(expr, register)
	expr.Close()
	return typ, err
}

// ExpressionToRegister moves the result of an expression into the given register.
func (state *State) ExpressionToRegister(root *expression.Expression, finalRegister *register.Register) (*types.Type, error) {
	if root.IsLeaf() {
		return state.TokenToRegister(root.Token, finalRegister)
	}

	// Resolve package access
	err := state.ResolveAccessors(root)

	if err != nil {
		return nil, err
	}

	// Save the temporary registers so we can easily free them later
	var temporaryRegisters []*register.Register

	// Sort by expression complexity so that we can
	// calculate the most complex expression first.
	// This reduces the number of registers required.
	root.SortByRegisterCount()

	if finalRegister != nil {
		root.Register = finalRegister

		// Assign final register to the left operands in the left tree
		left := root

		for len(left.Children) > 0 {
			left = left.Children[0]
			left.Register = finalRegister
		}
	}

	// Execute each operation starting from the bottom left
	err = root.EachOperation(func(sub *expression.Expression) error {
		if sub.IsFunctionCall {
			// Allocate a temporary register if necessary
			if sub.Register == nil && sub.Parent != nil {
				sub.Register = state.registers.General.FindFree()

				if sub.Register == nil {
					return errors.New(errors.ExceededMaxVariables)
				}

				_ = sub.Register.Use(sub)
				temporaryRegisters = append(temporaryRegisters, sub.Register)
			}

			return state.CallExpression(sub)
		}

		left := sub.Children[0]
		right := sub.Children[1]

		// Allocate a temporary register if necessary
		if left.Register == nil {
			left.Register = state.registers.General.FindFree()

			if left.Register == nil {
				return errors.New(errors.ExceededMaxVariables)
			}

			_ = left.Register.Use(sub)
			temporaryRegisters = append(temporaryRegisters, left.Register)
		}

		sub.Register = left.Register
		operator := sub.Token.Text()

		// Struct field access
		if operator == "." {
			variableName := left.Token.Text()
			variable := state.scopes.Get(variableName)

			if variable == nil {
				return errors.New(state.UnknownVariableError(variableName))
			}

			state.UseVariable(variable)
			fieldName := right.Token.Text()
			field := variable.Type.FieldByName(fieldName)

			if field == nil {
				return errors.New(&errors.UnknownField{Name: fieldName})
			}

			sub.Type = field.Type
			state.assembler.LoadRegister(sub.Register, variable.Register(), byte(field.Offset), byte(field.Type.Size))
			return nil
		}

		// Left operand
		if left.IsLeaf() {
			typ, err := state.TokenToRegister(left.Token, sub.Register)

			if err != nil {
				return err
			}

			left.Type = typ
		} else if sub.Register != left.Register {
			state.assembler.MoveRegisterRegister(sub.Register, left.Register)
			left.Register.Free()
		}

		if sub.Type == nil {
			sub.Type = left.Type
		}

		// Right operand is a leaf node
		if right.IsLeaf() {
			switch right.Token.Kind {
			case token.Identifier:
				variableName := right.Token.Text()
				variable := state.scopes.Get(variableName)

				if variable == nil {
					return errors.New(state.UnknownVariableError(variableName))
				}

				state.UseVariable(variable)
				right.Type = variable.Type
				return state.CalculateRegisterRegister(operator, sub.Register, variable.Register())

			case token.Number:
				right.Type = types.Int
				return state.CalculateRegisterNumber(operator, sub.Register, right)

			default:
				return fmt.Errorf("Invalid operand %s", right.Token)
			}
		}

		// Right operand is an expression
		return state.CalculateRegisterRegister(operator, sub.Register, right.Register)
	})

	if err != nil {
		return nil, err
	}

	// Free temporary registers
	for _, reg := range temporaryRegisters {
		reg.Free()
	}

	// Mark final register as used if it's not marked already
	if finalRegister != nil && finalRegister.IsFree() {
		_ = finalRegister.Use(root)
	}

	return root.Type, nil
}

// TokenToRegister moves a token into a register.
// It only works with identifiers, numbers and texts.
func (state *State) TokenToRegister(singleToken token.Token, register *register.Register) (*types.Type, error) {
	switch singleToken.Kind {
	case token.Identifier:
		variableName := singleToken.Text()
		variable := state.scopes.Get(variableName)

		if variable == nil {
			return nil, errors.New(state.UnknownVariableError(variableName))
		}

		state.UseVariable(variable)

		// Moving a variable into its own register is pointless
		if variable.Register() == register {
			return nil, nil
		}

		state.assembler.MoveRegisterRegister(register, variable.Register())
		return variable.Type, nil

	case token.Number:
		numberString := singleToken.Text()
		number, err := state.ParseInt(numberString)

		if err != nil {
			return nil, err
		}

		state.assembler.MoveRegisterNumber(register, uint64(number))
		return types.Int, nil

	case token.Text:
		address := state.assembler.AddString(singleToken.Text())
		state.assembler.MoveRegisterAddress(register, address)
		return types.Text, nil
	}

	return nil, errors.New(errors.NotImplemented)
}

// CalculateRegisterNumber performs an operation on a register and a number.
func (state *State) CalculateRegisterNumber(operation string, register *register.Register, operand *expression.Expression) error {
	number, err := state.ParseInt(operand.Token.Text())

	if err != nil {
		return err
	}

	switch operation {
	case "+":
		if number == 1 {
			state.assembler.IncreaseRegister(register)
			return nil
		}

		state.assembler.AddRegisterNumber(register, uint64(number))

	case "-":
		if number == 1 {
			state.assembler.DecreaseRegister(register)
			return nil
		}

		state.assembler.SubRegisterNumber(register, uint64(number))

	case "*":
		state.assembler.MulRegisterNumber(register, uint64(number))

	case "/":
		temporary := state.registers.General.FindFree()

		if temporary == nil {
			return errors.New(errors.ExceededMaxVariables)
		}

		temporary.ForceUse(operand)
		state.assembler.MoveRegisterNumber(temporary, uint64(number))
		err := state.CalculateRegisterRegister("/", register, temporary)

		if err != nil {
			return err
		}

		temporary.Free()

	default:
		return errors.New(errors.NotImplemented)
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

	case "/":
		rax := state.registers.All.ByName("rax")
		rdx := state.registers.All.ByName("rdx")

		if rax != registerTo {
			err := state.TryFreeRegister(rax)

			if err != nil {
				return err
			}

			state.assembler.MoveRegisterRegister(rax, registerTo)
		}

		err := state.TryFreeRegister(rdx)

		if err != nil {
			return err
		}

		state.assembler.SignExtendToDX(rax)
		state.assembler.DivRegister(registerFrom)
		state.assembler.MoveRegisterRegister(registerTo, rax)

	default:
		return errors.New(errors.NotImplemented)
	}

	return nil
}

// TryFreeRegister tries to free a register by moving its current user to another register.
func (state *State) TryFreeRegister(reg *register.Register) error {
	if reg.IsFree() {
		return nil
	}

	freeRegister := state.registers.General.FindFree()

	if freeRegister == nil {
		return errors.New(errors.ExceededMaxVariables)
	}

	state.assembler.MoveRegisterRegister(freeRegister, reg)
	variable, isVariable := reg.User().(*Variable)

	if !isVariable {
		return fmt.Errorf("User of register '%s' is not a variable", reg)
	}

	return variable.SetRegister(freeRegister)
}

// ResolveAccessors combines the children in the dot operator to a single function name.
func (state *State) ResolveAccessors(root *expression.Expression) error {
	for _, child := range root.Children {
		err := state.ResolveAccessors(child)

		if err != nil {
			return err
		}
	}

	return state.ResolveAccessor(root)
}

// ResolveAccessor combines the children in the dot operator to a single function name.
func (state *State) ResolveAccessor(root *expression.Expression) error {
	if root.Token.Text() != "." || !root.Children[1].IsFunctionCall {
		return nil
	}

	pkg := root.Children[0]
	pkgName := pkg.Token.Text()
	imp := state.function.File.imports[pkgName]

	if imp == nil {
		return errors.New(state.UnknownPackageError(pkgName))
	}

	atomic.AddInt32(&imp.Used, 1)
	newName := append(unsafe.StringToBytes(imp.Path), '.')
	newName = append(newName, root.Children[1].Token.Bytes...)
	root.Children[1].Token.Bytes = newName
	root.Replace(root.Children[1])
	return nil
}
