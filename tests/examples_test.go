package tests_test

import (
	"path/filepath"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/go/assert"
)

var examples = []testRun{
	{"hello", "", "Hello\n", 0},
}

func TestExamples(t *testing.T) {
	for _, test := range examples {
		test.Run(t, filepath.Join("..", "examples", test.Name))
	}
}

func BenchmarkExamples(b *testing.B) {
	for _, test := range examples {
		b.Run(test.Name, func(b *testing.B) {
			example := config.New(filepath.Join("..", "examples", test.Name))

			for b.Loop() {
				_, err := compiler.Compile(example)
				assert.Nil(b, err)
			}
		})
	}
}