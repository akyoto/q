package cpu

import "fmt"

// SystemRegister represents a system register.
type SystemRegister int8

// String returns the human readable name of the system register.
func (r SystemRegister) String() string {
	if r < 0 {
		return "sys?"
	}

	return fmt.Sprintf("sys%d", r)
}