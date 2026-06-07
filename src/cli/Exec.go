package cli

// Exec runs the command included in the first argument and returns the exit code.
func Exec(args []string) int {
	if len(args) == 0 {
		return invalid()
	}

	switch args[0] {
	case "asm":
		return asm(args[1:])

	case "build":
		return build(args[1:])

	case "files":
		return files(args[1:])

	case "help":
		return help()

	case "keywords":
		return keywords()

	case "run":
		return run(args[1:])

	case "ssa":
		return ssa(args[1:])

	case "version":
		return version()

	default:
		return run(args)
	}
}