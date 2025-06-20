package ssa_test

import (
	"os"
	"runtime/debug"
	"testing"

	"git.urbach.dev/go/assert"
)

// This benchmark compares the performance of fat structs and interfaces.
// It allocates `n` objects where `n` must be divisible by 2.
const n = 100

type FatStruct struct {
	Type byte
	A    int
	B    int
	C    int
	D    int
	E    int
	F    int
	G    int
}

type Instruction interface{}

type BinaryInstruction struct {
	A int
	B int
}

type OtherInstruction struct {
	C int
	D int
	E int
	F int
	G int
}

func TestMain(m *testing.M) {
	debug.SetGCPercent(-1)
	os.Exit(m.Run())
}

func BenchmarkFatStructRaw(b *testing.B) {
	for b.Loop() {
		entries := make([]FatStruct, 0, n)

		for i := range n {
			entries = append(entries, FatStruct{
				Type: byte(i % 2),
				A:    i,
				B:    i,
			})
		}

		count := 0

		for _, entry := range entries {
			switch entry.Type {
			case 0:
				count++
			case 1:
			}
		}

		assert.Equal(b, count, n/2)
	}
}

func BenchmarkFatStructPtr(b *testing.B) {
	for b.Loop() {
		entries := make([]*FatStruct, 0, n)

		for i := range n {
			entries = append(entries, &FatStruct{
				Type: byte(i % 2),
				A:    i,
				B:    i,
			})
		}

		count := 0

		for _, entry := range entries {
			switch entry.Type {
			case 0:
				count++
			case 1:
			}
		}

		assert.Equal(b, count, n/2)
	}
}

func BenchmarkInterfaceRaw(b *testing.B) {
	for b.Loop() {
		entries := make([]Instruction, 0, n)

		for i := range n {
			if i%2 == 0 {
				entries = append(entries, BinaryInstruction{
					A: i,
					B: i,
				})
			} else {
				entries = append(entries, OtherInstruction{
					C: i,
					D: i,
				})
			}
		}

		count := 0

		for _, entry := range entries {
			switch entry.(type) {
			case BinaryInstruction:
				count++
			case OtherInstruction:
			}
		}

		assert.Equal(b, count, n/2)
	}
}

func BenchmarkInterfacePtr(b *testing.B) {
	for b.Loop() {
		entries := make([]Instruction, 0, n)

		for i := range n {
			if i%2 == 0 {
				entries = append(entries, &BinaryInstruction{
					A: i,
					B: i,
				})
			} else {
				entries = append(entries, &OtherInstruction{
					C: i,
					D: i,
				})
			}
		}

		count := 0

		for _, entry := range entries {
			switch entry.(type) {
			case *BinaryInstruction:
				count++
			case *OtherInstruction:
			}
		}

		assert.Equal(b, count, n/2)
	}
}