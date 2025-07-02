package ssa2asm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/go/assert"
)

func TestHelloExample(t *testing.T) {
	b := build.New("../../examples/hello")
	systems := []build.OS{build.Linux, build.Mac, build.Windows}
	architectures := []build.Arch{build.ARM, build.X86}

	for _, os := range systems {
		b.OS = os

		for _, arch := range architectures {
			b.SetArch(arch)
			_, err := compiler.Compile(b)
			assert.Nil(t, err)
		}
	}
}