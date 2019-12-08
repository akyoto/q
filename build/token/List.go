package token

import "strings"

// List is a slice of tokens.
type List []Token

// String implements string serialization.
func (list List) String() string {
	builder := strings.Builder{}

	for _, t := range list {
		builder.WriteString(t.String())
	}

	return builder.String()
}
