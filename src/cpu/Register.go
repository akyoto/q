package cpu

import "fmt"

// Register represents the number of the register.
type Register int8

// String returns the human readable name of the register.
func (r Register) String() string {
	if r < 0 {
		return "r?"
	}

	return fmt.Sprintf("r%d", r)
}