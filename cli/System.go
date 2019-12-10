package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/akyoto/color"
	"github.com/akyoto/q/build"
	"github.com/akyoto/q/build/log"
)

// System shows system information.
func System() {
	const prefix = "%-19s"
	key := color.New(color.Faint).Sprint
	value := fmt.Sprint
	errorValue := color.RedString

	// Platform
	log.Info.Printf(key(prefix)+"%s\n", "Platform:", value(runtime.GOOS))

	// Architecture
	log.Info.Printf(key(prefix)+"%s\n", "Architecture:", value(runtime.GOARCH))

	// Go version
	log.Info.Printf(key(prefix)+"%s\n", "Go version:", value(runtime.Version()))

	// Working directory
	workingDirectory, err := os.Getwd()

	if err == nil {
		log.Info.Printf(key(prefix)+"%s\n", "Working directory:", value(workingDirectory))
	} else {
		log.Info.Printf(key(prefix)+"%s\n", "Working directory:", errorValue(err.Error()))
	}

	// Executable path
	executable, err := os.Executable()

	if err == nil {
		log.Info.Printf(key(prefix)+"%s\n", "Compiler path:", value(executable))
	} else {
		log.Info.Printf(key(prefix)+"%s\n", "Compiler path:", errorValue(err.Error()))
	}

	// Standard library path
	stdLib, err := build.FindStandardLibrary()

	if err == nil {
		log.Info.Printf(key(prefix)+"%s\n", "Standard library:", value(stdLib))
	} else {
		log.Info.Printf(key(prefix)+"%s\n", "Standard library:", errorValue(err.Error()))
	}

	// CPU model
	cpuInfoBytes, err := ioutil.ReadFile("/proc/cpuinfo")

	if err == nil {
		cpuInfo := string(cpuInfoBytes)
		modelNamePrefix := "model name"
		cpuInfo = cpuInfo[strings.Index(cpuInfo, modelNamePrefix)+len(modelNamePrefix):]
		cpuInfo = strings.TrimSpace(cpuInfo)
		cpuInfo = strings.TrimPrefix(cpuInfo, ":")
		cpuInfo = strings.TrimSpace(cpuInfo)
		cpuInfo = cpuInfo[:strings.Index(cpuInfo, "\n")]
		log.Info.Printf(key(prefix)+"%s\n", "CPU:", value(cpuInfo))
	}

	// CPU threads
	cpuCount := runtime.NumCPU()
	log.Info.Printf(key(prefix)+"%s\n", "CPU threads:", value(strconv.Itoa(cpuCount)))
}
