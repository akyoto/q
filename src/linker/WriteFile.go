package linker

import (
	"os"

	"git.urbach.dev/cli/q/src/core"
)

// WriteFile writes an executable file to disk.
func WriteFile(executable string, env *core.Environment) error {
	file, err := os.Create(executable)

	if err != nil {
		return err
	}

	Write(file, env)
	err = file.Chmod(0o755)

	if err != nil {
		return err
	}

	return file.Close()
}