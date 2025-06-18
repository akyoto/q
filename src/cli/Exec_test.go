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
	assert.Equal(t, cli.Exec([]string{"run"}), 0)
	assert.Equal(t, cli.Exec([]string{"help"}), 0)
}