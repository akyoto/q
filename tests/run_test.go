package tests_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/linker"
	"git.urbach.dev/go/assert"
)

type run struct {
	Name     string
	Input    string
	Output   string
	ExitCode int
}

// Run builds and runs the file to check if the output matches the expected output.
func (test *run) Run(t *testing.T, path string) {
	t.Run(test.Name, func(t *testing.T) {
		build := config.New(path)
		env, err := compiler.Compile(build)
		assert.Nil(t, err)

		if test.ExitCode == -1 {
			return
		}

		tmpDir := filepath.Join(os.TempDir(), "q", "tests")
		err = os.MkdirAll(tmpDir, 0o755)
		assert.Nil(t, err)

		executable := build.Executable()
		executable = filepath.Join(tmpDir, filepath.Base(executable))
		err = linker.WriteFile(executable, env)
		assert.Nil(t, err)

		stat, err := os.Stat(executable)
		assert.Nil(t, err)
		assert.True(t, stat.Size() > 0)

		cmd := exec.Command(executable)
		cmd.Stdin = strings.NewReader(test.Input)
		output, err := cmd.Output()
		exitCode := 0

		if err != nil {
			exitError, ok := err.(*exec.ExitError)

			if !ok {
				t.Fatal(exitError)
			}

			exitCode = exitError.ExitCode()
		}

		assert.Equal(t, exitCode, test.ExitCode)
		assert.DeepEqual(t, string(output), test.Output)
	})
}