package tests_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/linker"
	"git.urbach.dev/go/assert"
)

var examples = []struct {
	Name     string
	Input    string
	Output   string
	ExitCode int
}{
	{"hello", "", "Hello\n", 0},
}

func TestExamples(t *testing.T) {
	for _, test := range examples {
		directory := filepath.Join("..", "examples", test.Name)
		run(t, directory, test.Input, test.Output, test.ExitCode)
	}
}

func BenchmarkExamples(b *testing.B) {
	for _, test := range examples {
		b.Run(test.Name, func(b *testing.B) {
			example := build.New(filepath.Join("..", "examples", test.Name))

			for b.Loop() {
				_, err := compiler.Compile(example)
				assert.Nil(b, err)
			}
		})
	}
}

// run builds and runs the file to check if the output matches the expected output.
func run(t *testing.T, path string, input string, expectedOutput string, expectedExitCode int) {
	b := build.New(path)
	env, err := compiler.Compile(b)
	assert.Nil(t, err)

	tmpDir := filepath.Join(os.TempDir(), "q", "tests")
	err = os.MkdirAll(tmpDir, 0o755)
	assert.Nil(t, err)

	executable := b.Executable()
	executable = filepath.Join(tmpDir, filepath.Base(executable))
	err = linker.WriteFile(executable, b, env)
	assert.Nil(t, err)

	stat, err := os.Stat(executable)
	assert.Nil(t, err)
	assert.True(t, stat.Size() > 0)

	cmd := exec.Command(executable)
	cmd.Stdin = strings.NewReader(input)
	output, err := cmd.Output()
	exitCode := 0

	if err != nil {
		exitError, ok := err.(*exec.ExitError)

		if !ok {
			t.Fatal(exitError)
		}

		exitCode = exitError.ExitCode()
	}

	assert.Equal(t, exitCode, expectedExitCode)
	assert.DeepEqual(t, string(output), expectedOutput)
}