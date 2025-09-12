package tests_test

import (
	"path/filepath"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/go/assert"
)

var examples = []run{
	{"hello", "", "Hello\n", 0},
	{"factorial", "", "120", 0},
	{"fibonacci", "", "55", 0},
	{"fizzbuzz", "", "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz", 0},
	{"gcd", "", "21", 0},
	{"collatz", "", "12 6 3 10 5 16 8 4 2 1", 0},
	{"prime", "", "2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97", 0},
	{"point", "", "Point: 1, 2", 0},
	{"echo", "Echo", "Echo", 0},
	{"readfile", "", "", -1},
	{"raylib", "", "", -1},
	{"server", "", "", -1},
	{"shell", "", "", -1},
	{"thread", "", "", -1},
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