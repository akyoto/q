package build

import "github.com/akyoto/q/spec"

// Functions defines the builtin functions.
var Functions = map[string]*Function{
	"print": {
		Name: "print",
		Parameters: []spec.Variable{
			{
				Name: "text",
				Type: spec.Types["Text"],
			},
		},
	},
}