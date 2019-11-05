package main_test

import (
	"os"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build"
)

func TestCLI(t *testing.T) {
	// Cleanup
	defer os.Remove("examples/hello/hello")

	// Build
	os.Args = []string{"q", "build", "examples/hello"}
	build.Main()

	// Verify the file exists
	stat, err := os.Stat("examples/hello/hello")
	assert.Nil(t, err)
	assert.True(t, stat.Size() > 0)
}

func TestCLIHelp(t *testing.T) {
	build.Help()
}

func TestBuildErrors(t *testing.T) {
	tests := []struct {
		File          string
		ExpectedError string
	}{
		{"build/testdata/errors/missing-opening-bracket.q", "Missing opening bracket"},
		{"build/testdata/errors/missing-closing-bracket.q", "Missing closing bracket"},
		{"build/testdata/errors/unknown-function.q", "Unknown function"},
		{"build/testdata/errors/unknown-function-suggestion.q", "Unknown function 'prin', did you mean 'print'?"},
	}

	for _, test := range tests {
		file := build.NewFile(test.File, nil)
		defer file.Close()
		err := file.Compile()
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), test.ExpectedError)
	}
}
