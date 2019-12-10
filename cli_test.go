package main_test

import (
	"os"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/cli"
)

func TestCLI(t *testing.T) {
	type cliTest struct {
		Arguments        []string
		ExpectedExitCode int
	}

	tests := []cliTest{
		{[]string{"q"}, 2},
		{[]string{"q", "invalid"}, 2},
		{[]string{"q", "build", "non-existing-directory"}, 1},
		{[]string{"q", "build", "examples/hello/hello.q"}, 2},
	}

	for _, example := range examples {
		tests = append(tests, cliTest{[]string{"q", "build", "-v", "examples/" + example.Name}, 0})
		tests = append(tests, cliTest{[]string{"q", "build", "-v", "-O", "examples/" + example.Name}, 0})
	}

	for _, test := range tests {
		os.Args = test.Arguments
		exitCode := cli.Main()
		t.Log(test.Arguments)
		assert.Equal(t, exitCode, test.ExpectedExitCode)
	}

	stat, err := os.Stat("examples/hello/hello")
	assert.Nil(t, err)
	assert.True(t, stat.Size() > 0)
}
