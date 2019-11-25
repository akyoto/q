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
		SideEffects: 1,
	},
	"syscall": {
		Name:             "syscall",
		NoParameterCheck: true,
		SideEffects:      1,
	},
}
