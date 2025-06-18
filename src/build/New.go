package build

import "runtime"

// New creates a new build.
func New(files ...string) *Build {
	b := &Build{
		Files: files,
	}

	switch runtime.GOARCH {
	case "amd64":
		b.Arch = X86
	case "arm64":
		b.Arch = ARM
	}

	switch runtime.GOOS {
	case "linux":
		b.OS = Linux
	case "darwin":
		b.OS = Mac
	case "windows":
		b.OS = Windows
	}

	return b
}