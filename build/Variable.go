package build

import (
	"github.com/akyoto/q/build/instruction"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/spec"
	"github.com/akyoto/q/build/token"
)

// Variable represents both local variables and function parameters.
type Variable struct {
	Name       string
	Type       *spec.Type
	AliveUntil instruction.Position
	Position   token.Position
	Mutable    bool
	register   *register.Register
}

// Register returns the register the variable refers to.
func (variable *Variable) Register() *register.Register {
	return variable.register
}

// SetRegister binds the variable to a register.
func (variable *Variable) SetRegister(register *register.Register) error {
	variable.register = register
	return register.Use(variable)
}

// ForceSetRegister binds the variable to a register regardless whether it's used or not.
func (variable *Variable) ForceSetRegister(register *register.Register) {
	variable.register = register
	register.Free()
	_ = register.Use(variable)
}

// String returns the string representation.
func (variable *Variable) String() string {
	return variable.Name
}
