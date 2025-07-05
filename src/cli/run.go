package cli

import (
	"fmt"
	"os"
	"os/exec"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/linker"
)

// run builds and runs the executable.
func run(args []string) int {
	b, err := newBuildFromArgs(args)

	if err != nil {
		return exit(err)
	}

	result, err := compiler.Compile(b)

	if err != nil {
		return exit(err)
	}

	err = linker.WriteFile(b.Executable(), b, result)

	if err != nil {
		return exit(err)
	}

	cmd := exec.Command(b.Executable())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exit(err)
	}

	return 0
}