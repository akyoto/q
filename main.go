package main

import (
	"os"
)

func main() {
	if len(os.Args) == 1 {
		help()
		os.Exit(2)
	}

	inputFile := os.Args[len(os.Args)-1]
	compiler := NewCompiler()
	compiler.Compile(inputFile, "a.out")
}

func help() {
	os.Stderr.WriteString("Missing input file\n")
}
