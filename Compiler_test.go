package main

import (
	"os"
	"testing"

	"github.com/akyoto/assert"
)

func TestCompiler(t *testing.T) {
	defer os.Remove("a.out")
	compiler := NewCompiler()
	err := compiler.Compile("testdata/compiler-bench-1k.q", "a.out")
	assert.Nil(t, err)
}

func BenchmarkCompiler(b *testing.B) {
	defer os.Remove("a.out")
	compiler := NewCompiler()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = compiler.Compile("testdata/compiler-bench-1k.q", "a.out")
	}
}
