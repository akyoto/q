package memfile_test

import (
	"errors"
	"io/fs"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/linker"
	"git.urbach.dev/cli/q/src/memfile"
	"git.urbach.dev/go/assert"
)

func TestEmptyFile(t *testing.T) {
	file, err := memfile.New("")
	assert.Nil(t, err)
	assert.NotNil(t, file)

	err = memfile.Exec(file)
	assert.NotNil(t, err)

	err = file.Close()
	assert.True(t, errors.Is(err, fs.ErrClosed))
}

func TestHelloExample(t *testing.T) {
	file, err := memfile.New("")
	assert.Nil(t, err)
	assert.NotNil(t, file)

	b := config.New("../../examples/hello")
	env, err := compiler.Compile(b)
	assert.Nil(t, err)

	linker.Write(file, env)
	err = memfile.Exec(file)
	assert.Nil(t, err)

	err = file.Close()
	assert.True(t, errors.Is(err, fs.ErrClosed))
}