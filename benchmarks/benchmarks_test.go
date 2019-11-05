package build_test

import (
	"testing"

	"github.com/akyoto/q/build"
)

func BenchmarkCompiler(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		file := build.NewFile("print.q", nil)
		err := file.Compile()

		if err != nil {
			b.Fatal(err)
		}

		file.Close()
	}
}
