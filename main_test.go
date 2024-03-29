package main_test

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build"
	"github.com/akyoto/q/build/log"
)

func TestMain(m *testing.M) {
	log.Info.SetOutput(io.Discard)
	log.Error.SetOutput(io.Discard)
	os.Exit(m.Run())
}

// Run builds and runs the program to
// check if the output matches the expected output.
func Run(t *testing.T, path string, expectedOutput string, expectedExitCode int) {
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
				t.Fatal(err)
			}

			exitCode = exitError.ExitCode()
		}

		assert.Equal(t, exitCode, expectedExitCode)
		assert.DeepEqual(t, string(output), expectedOutput)
	})
}

// Check creates a build with a single file.
func Check(inputFile string) error {
	compiler, err := build.New(filepath.Dir(inputFile))

	if err != nil {
		return err
	}

	functions, structs, imports, errors := build.FindFunctionsInFile(inputFile, compiler.MainPackage, compiler.Environment)
	err = compiler.Environment.Import(compiler.MainPackage, functions, structs, imports, errors)

	if err != nil {
		return err
	}

	_, err = compiler.Compile()
	return err
}
