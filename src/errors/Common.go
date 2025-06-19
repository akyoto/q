package errors

var (
	ExpectedFunctionDefinition = &String{"Expected function definition"}
	ExpectedPackageName        = &String{"Expected package name"}
	InvalidDefinition          = &String{"Invalid definition"}
	MissingBlockStart          = &String{"Missing '{'"}
	MissingBlockEnd            = &String{"Missing '}'"}
	MissingGroupStart          = &String{"Missing '('"}
	MissingGroupEnd            = &String{"Missing ')'"}
	MissingParameter           = &String{"Missing parameter"}
	MissingType                = &String{"Missing type"}
	NoInputFiles               = &String{"No input files"}
)