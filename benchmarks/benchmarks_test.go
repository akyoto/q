package build_test

import (
	"testing"

	"github.com/akyoto/q/build"
)

func BenchmarkTokenizer(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		file := build.NewFile("print.q")
		err := file.Compile()

		if err != nil {
			b.Fatal(err)
		}

		file.Close()
	}
}
