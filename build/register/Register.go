package register

import (
	"fmt"

	"github.com/akyoto/color"
)

// Register represents a single CPU register.
type Register struct {
	Name   string
	usedBy fmt.Stringer
}

// Use marks the register as used by the given object.
func (register *Register) Use(obj fmt.Stringer) error {
	if obj == nil {
		panic("register.Use parameter cannot be nil")
	}

	if register.usedBy != nil {
		return &ErrAlreadyInUse{register, register.usedBy}
	}

	register.usedBy = obj
	return nil
}

// ForceUse marks the register as used by the given object and cannot fail.
func (register *Register) ForceUse(obj fmt.Stringer) {
	register.usedBy = obj
}

// User returns the user of the register.
func (register *Register) User() fmt.Stringer {
	return register.usedBy
}

// UserString returns the user as a string.
// If the user is nil, it returns "?".
func (register *Register) UserString() string {
	if register.usedBy != nil {
		return register.usedBy.String()
	}

	return "?"
}

// Free frees the register so that it can be used for new calculations.
func (register *Register) Free() {
	register.usedBy = nil
}

// IsFree returns true if the register is not in use.
func (register *Register) IsFree() bool {
	return register.usedBy == nil
}

// String returns a human-readable representation of the register.
func (register *Register) String() string {
	return register.StringWithUser(register.UserString())
}

// StringWithUser returns a human-readable representation of the register.
func (register *Register) StringWithUser(usedBy string) string {
	return fmt.Sprintf("%s%s%v", register.Name, color.New(color.Faint).Sprint("="), color.GreenString(usedBy))
}
