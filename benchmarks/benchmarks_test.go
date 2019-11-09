package build_test

import (
	"testing"

	"github.com/akyoto/q/build"
)

func BenchmarkCalls(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tmp, err := build.New("calls")

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
