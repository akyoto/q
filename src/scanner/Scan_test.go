package scanner_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/scanner"
	"git.urbach.dev/go/assert"
)

func TestNotExisting(t *testing.T) {
	b := build.New("_")
	_, err := scanner.Scan(b)
	assert.NotNil(t, err)
}

func TestMultiPlatform(t *testing.T) {
	b := build.New("testdata/platforms")
	_, err := scanner.Scan(b)
	assert.Nil(t, err)
}

func TestScanHelloExample(t *testing.T) {
	b := build.New("../../examples/hello")
	_, err := scanner.Scan(b)
	assert.Nil(t, err)
}