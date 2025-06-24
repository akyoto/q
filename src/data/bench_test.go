package data_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/data"
)

func BenchmarkFinalize(b *testing.B) {
	d := data.Data{}
	d.Insert("1", []byte("Beautiful is better than ugly."))
	d.Insert("2", []byte("Explicit is better than implicit."))
	d.Insert("3", []byte("Simple is better than complex."))
	d.Insert("4", []byte("Complex is better than complicated."))
	d.Insert("5", []byte("Flat is better than nested."))
	d.Insert("6", []byte("Sparse is better than dense."))

	for b.Loop() {
		d.Finalize()
	}
}