package global

import (
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
)

// Global variables that are useful in all packages.
var (
	Arch       string
	Executable string
	Library    string
	OS         string
	Root       string
)

// init is the very first thing that's executed.
// It disables the GC and initializes global variables.
func init() {
	debug.SetGCPercent(-1)
	OS = runtime.GOOS
	Arch = runtime.GOARCH

	var err error
	Executable, err = os.Executable()

	if err != nil {
		panic(err)
	}

	Executable, err = filepath.EvalSymlinks(Executable)

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