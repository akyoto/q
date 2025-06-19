package cli_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/cli"
	"git.urbach.dev/go/assert"
)

func TestExec(t *testing.T) {
	assert.Equal(t, cli.Exec(nil), 2)
	assert.Equal(t, cli.Exec([]string{"_"}), 2)
	assert.Equal(t, cli.Exec([]string{"build"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "--invalid-parameter"}), 2)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--invalid-parameter"}), 2)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--os", "linux"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--os", "mac"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--os", "windows"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--arch", "arm"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--arch", "x86"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello/hello.q", "--dry"}), 0)
	assert.Equal(t, cli.Exec([]string{"help"}), 0)
	assert.Equal(t, cli.Exec([]string{"run"}), 0)
}