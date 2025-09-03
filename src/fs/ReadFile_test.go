package fs_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/go/assert"
)

func TestReadFile(t *testing.T) {
	contents, err := fs.ReadFile("ReadFile_test.go")
	assert.Nil(t, err)
	assert.NotEqual(t, len(contents), 0)
}