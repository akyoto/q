package types

import "strings"

// Function transforms inputs to new outputs.
type Function struct {
	Input  []Type
	Output []Type
}

// Name returns the type name.
func (f *Function) Name() string {
	builder := strings.Builder{}
	builder.WriteString("(")

	for i, input := range f.Input {
		builder.WriteString(input.Name())

		if i != len(f.Input)-1 {
			builder.WriteString(", ")
		}
	}

	builder.WriteString(")")

	if len(f.Output) == 0 {
		return builder.String()
	}

	builder.WriteString(" -> (")

	for i, output := range f.Output {
		builder.WriteString(output.Name())

		if i != len(f.Output)-1 {
			builder.WriteString(", ")
		}
	}

	builder.WriteString(")")
	return builder.String()
}

// Size returns the total size in bytes.
func (f *Function) Size() int {
	return 8
}