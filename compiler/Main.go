package compiler

import (
	"os"
	"path/filepath"
	"strings"
)

// Main is the entry point for the CLI frontend.
func Main() {
	if len(os.Args) == 1 {
		Help()
		os.Exit(2)
	}

	inputFile := os.Args[len(os.Args)-1]
	compiler := New()
	inputFileBase := filepath.Base(inputFile)
	outputFile := strings.TrimSuffix(inputFileBase, filepath.Ext(inputFileBase))
	err := compiler.Compile(inputFile, outputFile)

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
