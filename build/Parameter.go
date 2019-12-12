package build

import (
	"github.com/akyoto/q/build/types"
)

// Parameter represents a function parameter.
type Parameter types.Field

// String returns the name of the parameter.
func (parameter *Parameter) String() string {
	return parameter.Name
}
