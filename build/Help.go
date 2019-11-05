package build

import "os"

func Help() {
	os.Stderr.WriteString("q build [-v] [directory]\n")
}
