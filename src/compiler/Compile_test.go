package compiler_test

import (
	"errors"
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/go/assert"
)

func TestNotExisting(t *testing.T) {
	b := build.New("_")
	_, err := compiler.Compile(b)
	assert.NotNil(t, err)
}

func TestNoInputFiles(t *testing.T) {
	b := build.New(".")
	_, err := compiler.Compile(b)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, compiler.NoInputFiles))
}

func TestHelloExample(t *testing.T) {
	b := build.New("../../examples/hello")
	_, err := compiler.Compile(b)
	assert.Nil(t, err)
}