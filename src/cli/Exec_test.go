package cli_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/cli"
	"git.urbach.dev/go/assert"
)

func TestExec(t *testing.T) {
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello/hello.q", "--dry"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--os", "linux", "--arch", "arm"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--os", "linux", "--arch", "x86"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--os", "mac", "--arch", "arm"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--os", "mac", "--arch", "x86"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--os", "windows", "--arch", "arm"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--os", "windows", "--arch", "x86"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--assembly"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--intermediate"}), 0)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--dry", "--verbose"}), 0)
	assert.Equal(t, cli.Exec([]string{"help"}), 0)
	assert.Equal(t, cli.Exec([]string{"run", "../../examples/hello"}), 0)
	assert.Equal(t, cli.Exec([]string{"../../tests/script.q"}), 0)
}

func TestExecErrors(t *testing.T) {
	assert.Equal(t, cli.Exec([]string{"build"}), 1)
	assert.Equal(t, cli.Exec([]string{"run"}), 1)
	assert.Equal(t, cli.Exec([]string{"_"}), 1)
}

func TestExecWrongParameters(t *testing.T) {
	assert.Equal(t, cli.Exec(nil), 2)
	assert.Equal(t, cli.Exec([]string{"build", "--invalid-parameter"}), 2)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--invalid-parameter"}), 2)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--os"}), 2)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--os", "invalid-os"}), 2)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--arch"}), 2)
	assert.Equal(t, cli.Exec([]string{"build", "../../examples/hello", "--arch", "invalid-arch"}), 2)
}