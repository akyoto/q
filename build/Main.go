package build

import (
	"os"
)

// Main is the entry point for the CLI frontend.
func Main() {
	var (
		verbose   = false
		directory = "."
	)

	if len(os.Args) < 2 {
		Help()
		os.Exit(2)
	}

	command := os.Args[1]

	if command != "build" {
		Help()
		os.Exit(2)
	}

	for i := 2; i < len(os.Args); i++ {
		argument := os.Args[i]

		switch argument {
		case "-v":
			verbose = true

		default:
			directory = argument
			stat, err := os.Stat(directory)

			if err != nil {
				os.Stderr.WriteString(err.Error() + "\n")
				os.Exit(2)
			}

			if !stat.IsDir() {
				os.Stderr.WriteString("Build path must be a directory\n")
				os.Exit(2)
			}
		}
	}

	build, err := New(directory)

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	build.Verbose = verbose
	err = build.Run()

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
