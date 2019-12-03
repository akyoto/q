package build_test

import (
	"testing"

	"github.com/akyoto/q/build"
)

func BenchmarkCompiler(b *testing.B) {
	directories := []string{
		"empty",
		"calc",
		"calls",
		"single-import",
	}

	for _, directory := range directories {
		directory := directory

		b.Run(directory, func(b *testing.B) {
			bench(b, directory)
		})
	}
}

func bench(b *testing.B, directory string) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tmp, err := build.New(directory)

		if err != nil {
			b.Fatal(err)
		}

		tmp.WriteExecutable = false
		err = tmp.Run()

		if err != nil {
			b.Fatal(err)
		}
	}
}
