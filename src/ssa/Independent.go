package ssa

// Independent values have no inputs.
type Independent struct{}

// Inputs returns nil because an independent value has no inputs.
func (_ *Independent) Inputs() []Value { return nil }

// Replace does nothing because an independent value has no inputs.
func (_ *Independent) Replace(Value, Value) {}