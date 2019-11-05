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
func assertOutput(t *testing.T, path string, expectedOutput string) {
	build, err := build.New(path)
	assert.Nil(t, err)
	assert.True(t, len(build.ExecutablePath) > 0)
	defer build.Close()
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
		assert.Nil(t, err)
		assert.DeepEqual(t, string(output), expectedOutput)
	})
}
