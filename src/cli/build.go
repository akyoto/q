package cli

import (
	"runtime"
	"strings"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/linker"
	"git.urbach.dev/cli/q/src/verbose"
)

// build parses the arguments and creates a build.
func build(args []string) int {
	b, err := newBuild(args)

	if err != nil {
		return exit(err)
	}

	env, err := compiler.Compile(b)

	if err != nil {
		return exit(err)
	}

	verbose.Show(env)

	if b.Dry {
		return 0
	}

	err = linker.WriteFile(b.Executable(), env)
	return exit(err)
}

// newBuild creates a new build with the given arguments.
func newBuild(args []string) (*config.Build, error) {
	build := config.New()

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--arch":
			i++

			if i >= len(args) {
				return nil, &ExpectedParameter{Parameter: "arch"}
			}

			switch args[i] {
			case "arm":
				build.Arch = config.ARM
			case "x86":
				build.Arch = config.X86
			default:
				return nil, &InvalidValue{Value: args[i], Parameter: "arch"}
			}

		case "--assembly", "-asm":
			build.ShowASM = true

		case "--dry":
			build.Dry = true

		case "--intermediate", "-ir":
			build.ShowIR = true

		case "--os":
			i++

			if i >= len(args) {
				return nil, &ExpectedParameter{Parameter: "os"}
			}

			switch args[i] {
			case "linux":
				build.OS = config.Linux
			case "mac":
				build.OS = config.Mac
			case "windows":
				build.OS = config.Windows
			default:
				return nil, &InvalidValue{Value: args[i], Parameter: "os"}
			}

		case "-v", "--verbose":
			build.SetVerbose(true)

		default:
			if strings.HasPrefix(args[i], "-") {
				return nil, &UnknownParameter{Parameter: args[i]}
			}

			build.Files = append(build.Files, args[i])
		}
	}

	if build.OS == config.UnknownOS {
		return nil, &InvalidValue{Value: runtime.GOOS, Parameter: "os"}
	}

	if build.Arch == config.UnknownArch {
		return nil, &InvalidValue{Value: runtime.GOARCH, Parameter: "arch"}
	}

	if len(build.Files) == 0 {
		build.Files = append(build.Files, ".")
	}

	return build, nil
}