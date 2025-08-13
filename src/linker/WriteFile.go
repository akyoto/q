package linker

import (
	"os"

	"git.urbach.dev/cli/q/src/core"
)

// WriteFile writes an executable file to disk.
func WriteFile(executable string, env *core.Environment) error {
	file, err := os.OpenFile(executable, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o755)

	if err != nil {
		return err
	}

	Write(file, env)
	return file.Close()
}