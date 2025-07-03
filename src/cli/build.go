package cli

import (
	"runtime"
	"strings"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/linker"
)

// _build parses the arguments and creates a build.
func _build(args []string) int {
	b, err := newBuildFromArgs(args)

	if err != nil {
		return exit(err)
	}

	result, err := compiler.Compile(b)

	if err != nil {
		return exit(err)
	}

	if b.Dry {
		return 0
	}

	err = linker.WriteFile(b.Executable(), b, result)
	return exit(err)
}

// newBuildFromArgs creates a new build with the given arguments.
func newBuildFromArgs(args []string) (*build.Build, error) {
	b := build.New()

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--arch":
			i++

			if i >= len(args) {
				return b, &ExpectedParameter{Parameter: "arch"}
			}

			switch args[i] {
			case "arm":
				b.Arch = build.ARM
			case "x86":
				b.Arch = build.X86
			default:
				return b, &InvalidValue{Value: args[i], Parameter: "arch"}
			}

		case "--dry":
			b.Dry = true

		case "--os":
			i++

			if i >= len(args) {
				return b, &ExpectedParameter{Parameter: "os"}
			}

			switch args[i] {
			case "linux":
				b.OS = build.Linux
			case "mac":
				b.OS = build.Mac
			case "windows":
				b.OS = build.Windows
			default:
				return b, &InvalidValue{Value: args[i], Parameter: "os"}
			}

		case "-v", "--verbose":
			b.ShowSSA = true

		default:
			if strings.HasPrefix(args[i], "-") {
				return b, &UnknownParameter{Parameter: args[i]}
			}

			b.Files = append(b.Files, args[i])
		}
	}

	if b.OS == build.UnknownOS {
		return b, &InvalidValue{Value: runtime.GOOS, Parameter: "os"}
	}

	if b.Arch == build.UnknownArch {
		return b, &InvalidValue{Value: runtime.GOARCH, Parameter: "arch"}
	}

	if len(b.Files) == 0 {
		b.Files = append(b.Files, ".")
	}

	return b, nil
}