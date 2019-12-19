package errors

var (
	ExceededMaxParameters       = &simple{"Exceeded maximum number of parameters per function", false}
	ExceededMaxVariables        = &simple{"Exceeded maximum limit of variables per function", false}
	ExpectedVariable            = &simple{"Expected variable on the left side of the assignment", false}
	InvalidExpression           = &simple{"Invalid expression", false}
	InvalidFunctionName         = &simple{"A function can not be named 'func' or 'fn'", false}
	InvalidInstruction          = &simple{"Invalid instruction", false}
	MissingAssignmentOperator   = &simple{"Missing assignment operator", false}
	MissingAssignmentExpression = &simple{"Missing assignment expression", false}
	MissingEndingNewline        = &simple{"Missing newline at the end of the file", false}
	MissingFunctionName         = &simple{"Expected function name before '('", false}
	MissingParameter            = &simple{"Missing parameter", false}
	MissingRange                = &simple{"Missing range expression in for loop", false}
	MissingRangeStart           = &simple{"Missing starting value in range expression", false}
	MissingRangeLimit           = &simple{"Missing upper limit in range expression", true}
	MissingReturnType           = &simple{"Missing function return type", false}
	MissingStructName           = &simple{"Missing struct name", false}
	NotImplemented              = &simple{"Not implemented", false}
	ParameterOpeningBracket     = &simple{"Missing opening bracket '(' after the function name", false}
	ReturnWithoutFunctionType   = &simple{"Returning a value in a function without a return type", false}
	EnsureWithoutFunctionType   = &simple{"Ensuring a value in a function without a return type", false}
	TopLevel                    = &simple{"Only function definitions are allowed at the top level", false}
	UnnecessaryNewlines         = &simple{"More than 2 successive empty lines", false}
)
