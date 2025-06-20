package global

import (
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
)

// Global variables that are useful in all packages.
var (
	Executable       string
	Library          string
	Root             string
	WorkingDirectory string
	HostOS           string
	HostArch         string
)

// init is the very first thing that's executed.
// It disables the GC and initializes global variables.
func init() {
	debug.SetGCPercent(-1)
	HostOS = runtime.GOOS
	HostArch = runtime.GOARCH

	var err error
	Executable, err = os.Executable()

	if err != nil {
		panic(err)
	}

	Executable, err = filepath.EvalSymlinks(Executable)

	if err != nil {
		panic(err)
	}

	WorkingDirectory, err = os.Getwd()

	if err != nil {
		panic(err)
	}

	Root = filepath.Dir(Executable)
	Library = filepath.Join(Root, "lib")
	stat, err := os.Stat(Library)

	if err != nil || !stat.IsDir() {
		findLibrary()
	}
}

// findLibrary tries to go up each directory from the working directory and check for the existence of a "lib" directory.
// This is needed for tests to work correctly.
func findLibrary() {
	dir := WorkingDirectory

	for {
		Library = filepath.Join(dir, "lib")
		stat, err := os.Stat(Library)

		if err == nil && stat.IsDir() {
			return
		}

		if dir == "/" {
			panic("standard library not found")
		}

		dir = filepath.Dir(dir)
	}
}