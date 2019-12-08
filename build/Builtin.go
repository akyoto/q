package build

const (
	BuiltinSyscall = "syscall"
	BuiltinPrint   = "print"
	BuiltinMemnum  = "memnum"
)

// BuiltinFunctions defines the builtin functions.
var BuiltinFunctions = map[string]*Function{
	BuiltinPrint: {
		Name: BuiltinPrint,
		Parameters: []*Parameter{
			{Name: "text"},
		},
		IsBuiltin:   true,
		SideEffects: 1,
	},
	BuiltinMemnum: {
		Name: BuiltinMemnum,
		Parameters: []*Parameter{
			{Name: "ptr"},
			{Name: "value"},
			{Name: "byteCount"},
		},
		IsBuiltin:   true,
		SideEffects: 1,
	},
	BuiltinSyscall: {
		Name: BuiltinSyscall,
		Parameters: []*Parameter{
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
