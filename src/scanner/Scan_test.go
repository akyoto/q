package scanner_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/scanner"
)

func TestScan(t *testing.T) {
	b := build.New("../../examples/hello")
	scanner.Scan(b)
}