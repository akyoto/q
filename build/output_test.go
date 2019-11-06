package build_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build"
	"github.com/akyoto/q/token"
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

// syntaxChecker creates a compiler that is purely used for syntax checks.
func syntaxChecker(t testing.TB, inputFile string) *build.Compiler {
	tmp := &build.Build{}
	contents, err := ioutil.ReadFile(inputFile)
	assert.Nil(t, err)
	tokens, processed := token.Tokenize(contents, []token.Token{})
	assert.Equal(t, processed, len(contents))
	return build.NewCompiler(tokens, tmp)
}
