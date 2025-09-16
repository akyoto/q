package config

import (
	"git.urbach.dev/cli/q/src/global"
)

// New creates a new build.
func New(files ...string) *Build {
	b := &Build{
		Files: files,
	}

	switch global.Arch {
	case "amd64":
		b.Arch = X86
	case "arm64":
		b.Arch = ARM
	}

	switch global.OS {
	case "linux":
		b.OS = Linux
	case "darwin":
		b.OS = Mac
	case "windows":
		b.OS = Windows
	}

	b.Optimize(true)
	b.Lint(true)
	return b
}