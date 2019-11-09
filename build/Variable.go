package build

import (
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/spec"
	"github.com/akyoto/q/token"
)

// Variable represents both local variables and function parameters.
type Variable struct {
	Name      string
	Type      *spec.Type
	Mutable   bool
	TimesUsed uint
	Position  token.Position
	Register  *register.Register
}

// BindRegister binds the variable to a register.
func (variable *Variable) BindRegister(register *register.Register) {
	variable.Register = register
	register.UsedBy = variable
}

// String returns the string representation.
func (variable *Variable) String() string {
	return variable.Name
}
