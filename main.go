package main

import (
	"os"

	"git.urbach.dev/cli/q/src/cli"
)

func main() {
	os.Exit(cli.Exec(os.Args[1:]))
}