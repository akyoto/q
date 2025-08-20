package cli

import "os"

// Exec runs the command included in the first argument and returns the exit code.
func Exec(args []string) int {
	if len(args) == 0 {
		return invalid()
	}

	switch args[0] {
	case "build":
		return build(args[1:])

	case "run":
		return run(args[1:])

	case "help":
		return help()

	case "version":
		return version()

	default:
		_, err := os.Stat(args[0])

		if err != nil {
			invalid()
		}

		return run(args)
	}
}