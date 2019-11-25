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
	Register   *register.Register
	Mutable    bool
}

// BindRegister binds the variable to a register.
func (variable *Variable) BindRegister(register *register.Register) {
	variable.Register = register
	register.Use(variable)
}

// String returns the string representation.
func (variable *Variable) String() string {
	return variable.Name
}
