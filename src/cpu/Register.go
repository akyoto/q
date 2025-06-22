package cpu

import "fmt"

// Register represents the number of the register.
type Register uint8

// String returns the human readable name of the register.
func (r Register) String() string {
	return fmt.Sprintf("r%d", r)
}