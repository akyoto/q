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
	err := compiler.Compile(inputFile, "a.out")

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func help() {
	os.Stderr.WriteString("Missing input file\n")
}
