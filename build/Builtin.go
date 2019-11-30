package build

import "github.com/akyoto/q/build/spec"

const (
	BuiltinSyscall = "syscall"
	BuiltinPrint   = "print"
)

// BuiltinFunctions defines the builtin functions.
var BuiltinFunctions = map[string]*Function{
	BuiltinPrint: {
		Name: BuiltinPrint,
		Parameters: []*Variable{
			{
				Name: "text",
				Type: spec.Types["Text"],
			},
		},
		IsBuiltin:   true,
		SideEffects: 1,
	},
	BuiltinSyscall: {
		Name: BuiltinSyscall,
		Parameters: []*Variable{
			{Name: "syscall number"},
			{Name: "param1"},
			{Name: "param2"},
			{Name: "param3"},
			{Name: "param4"},
			{Name: "param5"},
			{Name: "param6"},
		},
		NoParameterCheck: true,
		IsBuiltin:        true,
		SideEffects:      1,
	},
}
