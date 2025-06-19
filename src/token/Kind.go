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
	For                               // for
	If                                // if
	Import                            // import
	Loop                              // loop
	Return                            // return
	Switch                            // switch
	___END_KEYWORDS___                // </keywords>
)