package token

// Kind represents the type of token.
type Kind uint8

const (
	Invalid               Kind = iota // Invalid is an invalid token.
	EOF                               // EOF is the end of file.
	NewLine                           // NewLine is the newline character.
	Identifier                        // Identifier is a series of characters used to identify a variable or function.
	Number                            // Number is a series of numerical characters.
	Rune                              // Rune is a single unicode code point.
	String                            // String is an uninterpreted series of characters in the source code.
	Comment                           // Comment is a comment.
	Script                            // Script is a shebang line.
	GroupStart                        // (
	GroupEnd                          // )
	BlockStart                        // {
	BlockEnd                          // }
	ArrayStart                        // [
	ArrayEnd                          // ]
	ReturnType                        // ->
	___OPERATORS___                   // <operators>
	Add                               // +
	Sub                               // -
	Mul                               // *
	Div                               // /
	Mod                               // %
	And                               // &
	Or                                // |
	Xor                               // ^
	Shl                               // <<
	Shr                               // >>
	LogicalAnd                        // &&
	LogicalOr                         // ||
	Define                            // :=
	Dot                               // .
	Range                             // ..
	Call                              // x()
	Array                             // [x]
	Struct                            // x{}
	Separator                         // ,
	___ASSIGNMENTS___                 // <assignments>
	Assign                            // =
	AddAssign                         // +=
	SubAssign                         // -=
	MulAssign                         // *=
	DivAssign                         // /=
	ModAssign                         // %=
	AndAssign                         // &=
	OrAssign                          // |=
	XorAssign                         // ^=
	ShlAssign                         // <<=
	ShrAssign                         // >>=
	FieldAssign                       // :
	___END_ASSIGNMENTS___             // </assignments>
	___COMPARISONS___                 // <comparisons>
	Equal                             // ==
	NotEqual                          // !=
	Less                              // <
	Greater                           // >
	LessEqual                         // <=
	GreaterEqual                      // >=
	___END_COMPARISONS___             // </comparisons>
	___UNARY___                       // <unary>
	Not                               // ! (unary)
	Negate                            // - (unary)
	___END_UNARY___                   // </unary>
	___END_OPERATORS___               // </operators>
	___KEYWORDS___                    // <keywords>
	Assert                            // assert
	Const                             // const
	Else                              // else
	Extern                            // extern
	If                                // if
	Import                            // import
	Loop                              // loop
	Return                            // return
	Switch                            // switch
	___END_KEYWORDS___                // </keywords>
)

// IsAssignment returns true if the token is an assignment operator.
func (k Kind) IsAssignment() bool {
	return k > ___ASSIGNMENTS___ && k < ___END_ASSIGNMENTS___
}

// IsComparison returns true if the token is a comparison operator.
func (k Kind) IsComparison() bool {
	return k > ___COMPARISONS___ && k < ___END_COMPARISONS___
}

// IsExpressionStart returns true if the token starts an expression.
func (k Kind) IsExpressionStart() bool {
	return k == GroupStart || k == ArrayStart || k == BlockStart
}

// IsKeyword returns true if the token is a keyword.
func (k Kind) IsKeyword() bool {
	return k > ___KEYWORDS___ && k < ___END_KEYWORDS___
}

// IsNumeric returns true if the token is a number or rune.
func (k Kind) IsNumeric() bool {
	return k == Number || k == Rune
}

// IsOperator returns true if the token is an operator.
func (k Kind) IsOperator() bool {
	return k > ___OPERATORS___ && k < ___END_OPERATORS___
}

// IsUnaryOperator returns true if the token is a unary operator.
func (k Kind) IsUnaryOperator() bool {
	return k > ___UNARY___ && k < ___END_UNARY___
}

// String returns a human-readable representation of the token kind.
func (k Kind) String() string {
	switch k {
	case Struct:
		return "$"
	case Dot:
		return "."
	case Call:
		return "λ"
	case Array:
		return "@"
	case Negate:
		return "-"
	case Not:
		return "!"
	case Mul:
		return "*"
	case Div:
		return "/"
	case Mod:
		return "%"
	case Add:
		return "+"
	case Sub:
		return "-"
	case Shr:
		return ">>"
	case Shl:
		return "<<"
	case And:
		return "&"
	case Xor:
		return "^"
	case Or:
		return "|"
	case Greater:
		return ">"
	case Less:
		return "<"
	case GreaterEqual:
		return ">="
	case LessEqual:
		return "<="
	case Equal:
		return "=="
	case NotEqual:
		return "!="
	case LogicalAnd:
		return "&&"
	case LogicalOr:
		return "||"
	case Range:
		return ".."
	case Separator:
		return ","
	case Assign:
		return "="
	case Define:
		return ":="
	case AddAssign:
		return "+="
	case SubAssign:
		return "-="
	case MulAssign:
		return "*="
	case DivAssign:
		return "/="
	case ModAssign:
		return "%="
	case AndAssign:
		return "&="
	case OrAssign:
		return "|="
	case XorAssign:
		return "^="
	case ShrAssign:
		return ">>="
	case ShlAssign:
		return "<<="
	case FieldAssign:
		return ":"
	default:
		return ""
	}
}