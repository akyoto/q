package tests_test

import (
	"path/filepath"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/go/assert"
)

var examples = []run{
	{"hello", nil, "", "Hello\n", 0},
	{"factorial", nil, "", "120", 0},
	{"fibonacci", nil, "", "55", 0},
	{"fizzbuzz", nil, "", "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz", 0},
	{"gcd", nil, "", "21", 0},
	{"collatz", nil, "", "12 6 3 10 5 16 8 4 2 1", 0},
	{"prime", nil, "", "2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97", 0},
	{"point", nil, "", "Point: 1, 2", 0},
	{"echo", nil, "Echo", "Echo", 0},
	{"clock", nil, "", "", -1},
	{"raylib", nil, "", "", -1},
	{"readfile", nil, "", "", -1},
	{"server", nil, "", "", -1},
	{"shell", nil, "", "", -1},
	{"thread", nil, "", "", -1},
}

func TestExamples(t *testing.T) {
	for _, test := range examples {
		build := config.New(filepath.Join("..", "examples", test.Name))
		test.Run(t, build)
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