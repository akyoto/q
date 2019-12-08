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

	// Send and receive
	"->": {"->", 3, true},
	"<-": {"->", 3, true},

	// Logical OR
	"||": {"||", 4, true},

	// Logical AND
	"&&": {"&&", 5, true},

	// Comparison
	"==": {"==", 6, false},
	"!=": {"!=", 6, false},
	"<=": {"<=", 6, true},
	">=": {">=", 6, true},

	"<": {"<", 7, true},
	">": {">", 7, true},

	// Arithmetic operations
	"+": {"+", 8, false},
	"-": {"-", 8, false},

	"*": {"*", 9, false},
	"/": {"/", 9, true},
	"%": {"%", 9, true},

	// Package and field access
	".": {".", 10, true},
}
