package cli_test

import (
	"os"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/cli"
)

func TestCLIMain(t *testing.T) {
	defer os.Remove("../examples/hello/hello")

	os.Args = []string{"q", "build", "../examples/hello"}
	cli.Main()

	stat, err := os.Stat("../examples/hello/hello")
	assert.Nil(t, err)
	assert.True(t, stat.Size() > 0)
}
