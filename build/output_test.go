package build_test

import (
	"os"
	"os/exec"
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

// syntaxCheck treats the file as if it as function body and returns syntax errors.
func syntaxCheck(inputFile string) error {
	file := build.NewFile(inputFile)
	err := file.Tokenize()

	if err != nil {
		return err
	}

	function := &build.Function{
		File:       file,
		TokenStart: 0,
		TokenEnd:   len(file.Tokens()),
	}

	_, err = build.Compile(function, build.NewEnvironment(), false)
	return err
}
