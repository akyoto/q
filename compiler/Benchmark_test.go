package compiler_test

import (
	"testing"

	"github.com/akyoto/q/compiler"
)

func BenchmarkCompiler(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c := compiler.New()
		c.WriteExecutable = false
		_ = c.Compile("testdata/compiler-bench-1k.q", "")
		c.Close()
	}
}
