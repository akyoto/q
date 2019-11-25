package errors

var (
	NotImplemented              = &Base{"Not implemented", false}
	InvalidInstruction          = &Base{"Invalid instruction", false}
	MissingParameter            = &Base{"Missing parameter", false}
	MissingFunctionName         = &Base{"Expected function name before '('", false}
	ExpectedVariable            = &Base{"Expected variable on the left side of the assignment", false}
	InvalidFunctionName         = &Base{"A function can not be named 'func' or 'fn'", false}
	ParameterOpeningBracket     = &Base{"Missing opening bracket '(' after the function name", false}
	TopLevel                    = &Base{"Only function definitions are allowed at the top level", false}
	MissingRange                = &Base{"Missing range expression in for loop", false}
	MissingRangeLimit           = &Base{"Missing upper limit in range expression", true}
	MissingAssignmentOperator   = &Base{"Missing assignment operator", false}
	MissingAssignmentExpression = &Base{"Missing assignment expression", false}
)
