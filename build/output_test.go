package build_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build"
)

// assertOutput builds and runs the program to
// check if the output matches the expected output.
func assertOutput(t *testing.T, path string, expectedOutput string, expectedExitCode int) {
	build, err := build.New(path)
	assert.Nil(t, err)
	assert.True(t, len(build.ExecutablePath) > 0)
	defer os.Remove(build.ExecutablePath)

	t.Run("Compile", func(t *testing.T) {
		err = build.Run()
		assert.Nil(t, err)

		stat, err := os.Stat(build.ExecutablePath)
		assert.Nil(t, err)
		assert.True(t, stat.Size() > 0)
	})

	t.Run("Output", func(t *testing.T) {
		cmd := exec.Command(build.ExecutablePath)
		output, err := cmd.Output()
		exitCode := 0

		if err != nil {
			exitError, ok := err.(*exec.ExitError)

			if !ok {
				panic(err)
			}

			exitCode = exitError.ExitCode()
		}

		assert.Equal(t, exitCode, expectedExitCode)
		assert.DeepEqual(t, string(output), expectedOutput)
	})
}

// check creates a build with a single file.
func check(inputFile string) error {
	compiler, err := build.New(filepath.Dir(inputFile))

	if err != nil {
		return err
	}

	files := make(chan *build.File, 1)
	files <- build.NewFile(inputFile)
	close(files)

	functions, imports, errors := build.FindFunctions(files, compiler.Environment)
	err = compiler.Environment.Import("", functions, imports, make(chan error), errors)

	if err != nil {
		return err
	}

	return compiler.Compile()
}
