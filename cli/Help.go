package cli

import (
	"github.com/akyoto/color"
	"github.com/akyoto/q/build/log"
)

// Help shows the command line argument usage.
func Help() {
	log.Error.Println(color.YellowString("# Commands"))
	log.Error.Println("")
	log.Error.Println("q build", color.New(color.Faint).Sprint("[directory]"))
	log.Error.Println("q system")
	log.Error.Println("")
	log.Error.Println(color.YellowString("## build"))
	log.Error.Println("")
	log.Error.Println("Builds an executable from the source files in the directory.")
	log.Error.Println("")
	log.Error.Println("-v --verbose  Shows assembly output.")
	log.Error.Println("-O --optimize Optimizes for performance.")
	log.Error.Println("")
	log.Error.Println(color.YellowString("## system"))
	log.Error.Println("")
	log.Error.Println("Displays information about the system.")
}
