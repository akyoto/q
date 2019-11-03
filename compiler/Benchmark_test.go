package compiler_test

import (
	"testing"

	"github.com/akyoto/q/compiler"
)

func BenchmarkCompiler(b *testing.B) {
	c := compiler.New()
	c.WriteExecutable = false

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = c.Compile("testdata/compiler-bench-1k.q", "")
	}
}
