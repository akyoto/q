package register

import "fmt"

// Register represents a single CPU register.
type Register struct {
	Name   string
	UsedBy fmt.Stringer
}
