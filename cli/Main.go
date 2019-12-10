package cli

import (
	"os"

	"github.com/akyoto/q/build"
	"github.com/akyoto/q/build/log"
)

// Main is the entry point for the CLI frontend.
// It returns the exit code of the compiler.
// We never call os.Exit directly here because it's bad for testing.
func Main() int {
	var (
		verbose   = false
		optimize  = false
		directory = "."
	)

	if len(os.Args) < 2 {
		Help()
		return 2
	}

	command := os.Args[1]

	if command == "help" || command == "--help" {
		Help()
		return 0
	}

	if command == "system" {
		System()
		return 0
	}

	if command != "build" {
		Help()
		return 2
	}

	for i := 2; i < len(os.Args); i++ {
		argument := os.Args[i]

		switch argument {
		case "-v", "--verbose":
			verbose = true

		case "-O", "--optimize":
			optimize = true

		default:
			directory = argument
			stat, err := os.Stat(directory)

			if err != nil {
				log.Error.Println(err)
				return 1
			}

			if !stat.IsDir() {
				log.Error.Println("Build path must be a directory")
				return 2
			}
		}
	}

	b, err := build.New(directory)

	if err != nil {
		log.Error.Println(err)
		return 1
	}

	b.Verbose = verbose
	b.Optimize = optimize
	err = b.Run()

	if err != nil {
		log.Error.Println(err)
		return 1
	}

	return 0
}
