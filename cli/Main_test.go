package cli_test

import (
	"os"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/cli"
)

func TestCLI(t *testing.T) {
	tests := []struct {
		Arguments        []string
		ExpectedExitCode int
	}{
		{[]string{"q"}, 2},
		{[]string{"q", "invalid"}, 2},
		{[]string{"q", "build", "non-existing-directory"}, 1},
		{[]string{"q", "build", "../examples/hello/hello.q"}, 2},
		{[]string{"q", "build", "../examples/hello"}, 0},
		{[]string{"q", "build", "-v", "../examples/hello"}, 0},
		// {[]string{"q", "build", "../build/testdata"}, 1},
	}

	for _, test := range tests {
		os.Args = test.Arguments
		exitCode := cli.Main()
		assert.Equal(t, exitCode, test.ExpectedExitCode)
	}

	stat, err := os.Stat("../examples/hello/hello")
	assert.Nil(t, err)
	assert.True(t, stat.Size() > 0)
}
