package global

import (
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
)

// Global variables that are useful in all packages.
var (
	Arch    string
	Library string
	OS      string
)

// init is the very first thing that's executed.
// It disables the GC and initializes global variables.
func init() {
	debug.SetGCPercent(-1)
	OS = runtime.GOOS
	Arch = runtime.GOARCH

	executable, err := os.Executable()

	if err != nil {
		panic(err)
	}

	executable, err = filepath.EvalSymlinks(executable)

	if err != nil {
		panic(err)
	}

	root := filepath.Dir(executable)
	Library = filepath.Join(root, "lib")
	stat, err := os.Stat(Library)

	if err != nil || !stat.IsDir() {
		findLibrary()
	}
}