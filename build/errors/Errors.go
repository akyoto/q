package errors

var (
	ExceededMaxParameters       = &Base{"Exceeded maximum number of parameters per function", false}
	ExceededMaxVariables        = &Base{"Exceeded maximum limit of variables per function", false}
	ExpectedVariable            = &Base{"Expected variable on the left side of the assignment", false}
	InvalidExpression           = &Base{"Invalid expression", false}
	InvalidFunctionName         = &Base{"A function can not be named 'func' or 'fn'", false}
	InvalidInstruction          = &Base{"Invalid instruction", false}
	MissingAssignmentOperator   = &Base{"Missing assignment operator", false}
	MissingAssignmentExpression = &Base{"Missing assignment expression", false}
	MissingEndingNewline        = &Base{"Missing newline at the end of the file", false}
	MissingFunctionName         = &Base{"Expected function name before '('", false}
	MissingParameter            = &Base{"Missing parameter", false}
	MissingRange                = &Base{"Missing range expression in for loop", false}
	MissingRangeStart           = &Base{"Missing starting value in range expression", false}
	MissingRangeLimit           = &Base{"Missing upper limit in range expression", true}
	MissingType                 = &Base{"Missing type", false}
	NotImplemented              = &Base{"Not implemented", false}
	ParameterOpeningBracket     = &Base{"Missing opening bracket '(' after the function name", false}
	TopLevel                    = &Base{"Only function definitions are allowed at the top level", false}
	UnnecessaryNewlines         = &Base{"More than 2 successive empty lines", false}
)
