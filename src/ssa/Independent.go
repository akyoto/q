package ssa

// Independent values have no inputs.
type Independent struct{}

// Inputs returns nil because an independent value has no inputs.
func (*Independent) Inputs() []Value { return nil }

// Replace does nothing because an independent value has no inputs.
func (*Independent) Replace(Value, Value) {}