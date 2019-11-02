package main

import (
	"testing"

	"github.com/akyoto/assert"
)

func TestCompiler(t *testing.T) {
	compiler := NewCompiler()
	err := compiler.Compile("testdata/compiler-bench-1k.zen", "a.out")
	assert.Nil(t, err)
}

func BenchmarkCompiler(b *testing.B) {
	compiler := NewCompiler()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = compiler.Compile("testdata/compiler-bench-1k.zen", "a.out")
	}
}
