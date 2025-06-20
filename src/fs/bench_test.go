package fs_test

import (
	"os"
	"testing"

	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/go/assert"
)

func BenchmarkReadDir(b *testing.B) {
	for b.Loop() {
		files, err := os.ReadDir(".")
		assert.Nil(b, err)

		for _, file := range files {
			func(string) {}(file.Name())
		}
	}
}

func BenchmarkReaddirnames(b *testing.B) {
	for b.Loop() {
		f, err := os.Open(".")
		assert.Nil(b, err)
		files, err := f.Readdirnames(0)
		assert.Nil(b, err)

		for _, file := range files {
			func(string) {}(file)
		}

		assert.Nil(b, f.Close())
	}
}

func BenchmarkWalk(b *testing.B) {
	for b.Loop() {
		err := fs.Walk(".", func(string) {})
		assert.Nil(b, err)
	}
}