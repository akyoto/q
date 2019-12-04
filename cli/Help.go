package cli

import "github.com/akyoto/q/build/log"

// Help shows the command line argument usage.
func Help() {
	log.Error.Println("q build [-v] [-O] [directory]")
}
