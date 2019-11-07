package build

import "os"

// Help shows the command line argument usage.
func Help() {
	os.Stderr.WriteString("q build [-v] [directory]\n")
}
