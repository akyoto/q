package spec

// Functions defines the builtin functions.
var Functions = map[string]*Function{
	"print": {
		Name: "print",
		Parameters: []Variable{
			{
				Name: "text",
				Type: Types["Text"],
			},
		},
	},
}
