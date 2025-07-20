package compiler_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/go/assert"
)

func TestNotExisting(t *testing.T) {
	b := config.New("_")
	_, err := compiler.Compile(b)
	assert.NotNil(t, err)
}

func TestHelloExample(t *testing.T) {
	b := config.New("../../examples/hello")
	_, err := compiler.Compile(b)
	assert.Nil(t, err)
}

func TestHelloExampleVerbose(t *testing.T) {
	b := config.New("../../examples/hello")
	b.ShowSSA = true
	_, err := compiler.Compile(b)
	assert.Nil(t, err)
}