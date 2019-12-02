package spec

// Operators defines the operators used in the language.
// The number corresponds to the operator priority and can not be zero.
var Operators = map[string]*Operator{
	// Parameters
	// ",": {",", 1, true},

	// Assignment
	"=":   {"=", 2, true},
	"+=":  {"+=", 2, true},
	"-=":  {"-=", 2, true},
	"*=":  {"*=", 2, true},
	"/=":  {"/=", 2, true},
	">>=": {">>=", 2, true},
	"<<=": {"<<=", 2, true},

	// Logical OR
	"||": {"||", 3, true},

	// Logical AND
	"&&": {"&&", 4, true},

	// Comparison
	"==": {"==", 5, false},
	"!=": {"!=", 5, false},
	"<=": {"<=", 5, true},
	">=": {">=", 5, true},

	"<": {"<", 6, true},
	">": {">", 6, true},

	// Arithmetic operations
	"+": {"+", 7, false},
	"-": {"-", 7, false},

	"*": {"*", 8, false},
	"/": {"/", 8, true},
	"%": {"%", 8, true},

	// Package and field access
	".": {".", 9, true},
}
