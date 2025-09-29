package data_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/data"
)

func BenchmarkFinalize(b *testing.B) {
	d := data.Data{}
	d.SetImmutable("1", []byte("Beautiful is better than ugly."))
	d.SetImmutable("2", []byte("Explicit is better than implicit."))
	d.SetImmutable("3", []byte("Simple is better than complex."))
	d.SetImmutable("4", []byte("Complex is better than complicated."))
	d.SetImmutable("5", []byte("Flat is better than nested."))
	d.SetImmutable("6", []byte("Sparse is better than dense."))

	for b.Loop() {
		d.Finalize()
	}
}