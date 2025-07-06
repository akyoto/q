package memfile_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/memfile"
	"git.urbach.dev/go/assert"
)

func TestNew(t *testing.T) {
	file, err := memfile.New("")
	assert.Nil(t, err)
	assert.NotNil(t, file)
	// memfile.Exec can't be tested because it would replace the test executable
}