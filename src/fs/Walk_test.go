package fs_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/go/assert"
)

func TestWalk(t *testing.T) {
	var files []string

	err := fs.Walk(".", func(file string) {
		files = append(files, file)
	})

	assert.Nil(t, err)
	assert.Contains(t, files, "Walk_test.go")
}

func TestWalkNotDirectory(t *testing.T) {
	err := fs.Walk("Walk_test.go", func(file string) {})
	assert.NotNil(t, err)
}

func TestWalkNotExisting(t *testing.T) {
	err := fs.Walk("_", func(file string) {})
	assert.NotNil(t, err)
}