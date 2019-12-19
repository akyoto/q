package spec

// Operators defines the operators used in the language.
// The number corresponds to the operator priority and can not be zero.
var Operators = map[string]*Operator{
	// Parameters
	// ",": {",", 1, true, false},

	// Assignment
	"=":   {"=", 2, true, false},
	"+=":  {"+=", 2, true, false},
	"-=":  {"-=", 2, true, false},
	"*=":  {"*=", 2, true, false},
	"/=":  {"/=", 2, true, false},
	">>=": {">>=", 2, true, false},
	"<<=": {"<<=", 2, true, false},

	// Send and receive
	"->": {"->", 3, true, false},
	"<-": {"->", 3, true, false},

	// Logical OR
	"||": {"||", 4, true, false},

	// Logical AND
	"&&": {"&&", 5, true, false},

	// Comparison
	"==": {"==", 6, false, true},
	"!=": {"!=", 6, false, true},
	"<=": {"<=", 6, true, true},
	">=": {">=", 6, true, true},

	"<": {"<", 7, true, true},
	">": {">", 7, true, true},

	// Arithmetic operations
	"+": {"+", 8, false, false},
	"-": {"-", 8, false, false},

	"*": {"*", 9, false, false},
	"/": {"/", 9, true, false},
	"%": {"%", 9, true, false},

	// Package and field access
	".": {".", 10, true, false},
}
