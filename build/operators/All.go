package operators

// All defines the operators used in the language.
// The number corresponds to the operator priority and can not be zero.
var All = map[string]*Operator{
	// Parameters
	// ",": {",", 1, Default, true},

	// Assignment
	"=":   {"=", 2, Assignment, true},
	"+=":  {"+=", 2, Assignment, true},
	"-=":  {"-=", 2, Assignment, true},
	"*=":  {"*=", 2, Assignment, true},
	"/=":  {"/=", 2, Assignment, true},
	">>=": {">>=", 2, Assignment, true},
	"<<=": {"<<=", 2, Assignment, true},

	// Send and receive
	"->": {"->", 3, Default, true},
	"<-": {"->", 3, Default, true},

	// Logical OR
	"||": {"||", 4, Default, true},

	// Logical AND
	"&&": {"&&", 5, Default, true},

	// Comparison
	"==": {"==", 6, Comparison, false},
	"!=": {"!=", 6, Comparison, false},
	"<=": {"<=", 6, Comparison, true},
	">=": {">=", 6, Comparison, true},

	"<": {"<", 7, Comparison, true},
	">": {">", 7, Comparison, true},

	// Arithmetic operations
	"+": {"+", 8, Default, false},
	"-": {"-", 8, Default, false},

	"*": {"*", 9, Default, false},
	"/": {"/", 9, Default, true},
	"%": {"%", 9, Default, true},

	// Package and field access
	".": {".", 10, Default, true},
}
