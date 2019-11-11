package build

import "github.com/akyoto/q/build/spec"

// BuiltinFunctions defines the builtin functions.
var BuiltinFunctions = map[string]*Function{
	"print": {
		Name: "print",
		Parameters: []Variable{
			{
				Name: "text",
				Type: spec.Types["Text"],
			},
		},
	},
	"syscall": {
		Name:             "syscall",
		NoParameterCheck: true,
	},
}
