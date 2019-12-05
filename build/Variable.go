package build

import (
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/spec"
	"github.com/akyoto/q/build/token"
)

// Variable represents both local variables and function parameters.
type Variable struct {
	Name           string
	Type           *spec.Type
	AliveUntil     token.Position
	Position       token.Position
	LastAssign     token.Position
	LastAssignUsed bool
	Used           bool
	Mutable        bool
	register       *register.Register
}

// Register returns the register the variable refers to.
func (variable *Variable) Register() *register.Register {
	return variable.register
}

// SetRegister binds the variable to a register.
func (variable *Variable) SetRegister(register *register.Register) error {
	if variable.register != nil {
		variable.register.Free()
	}

	variable.register = register
	return register.Use(variable)
}

// ForceSetRegister binds the variable to a register regardless whether it's used or not.
func (variable *Variable) ForceSetRegister(register *register.Register) {
	_ = variable.SetRegister(register)
}

// String returns the string representation.
func (variable *Variable) String() string {
	return variable.Name
}
