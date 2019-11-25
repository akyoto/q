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
		HasSideEffects: true,
	},
	"syscall": {
		Name:             "syscall",
		NoParameterCheck: true,
		HasSideEffects:   true,
	},
}
