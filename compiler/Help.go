package compiler

import "os"

func Help() {
	os.Stderr.WriteString("Missing input file\n")
}
