package main

import (
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		help()
		os.Exit(2)
	}

	inputFile := os.Args[len(os.Args)-1]
	compiler := NewCompiler()
	inputFileBase := filepath.Base(inputFile)
	outputFile := strings.TrimSuffix(inputFileBase, filepath.Ext(inputFileBase))
	err := compiler.Compile(inputFile, outputFile)

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func help() {
	os.Stderr.WriteString("Missing input file\n")
}
