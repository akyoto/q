package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
	key := log.Faint.Sprint
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
	cpuModel := cpuModelName()

	if cpuModel != "" {
		log.Info.Printf(key(prefix)+"%s\n", "CPU:", value(cpuModel))
	}

	// CPU threads
	cpuCount := runtime.NumCPU()
	log.Info.Printf(key(prefix)+"%s\n", "CPU threads:", value(strconv.Itoa(cpuCount)))
}

// cpuModelName returns the model name of the CPU.
func cpuModelName() string {
	// Via "/proc/cpuinfo" (Linux)
	cpuInfoBytes, err := ioutil.ReadFile("/proc/cpuinfo")

	if err == nil {
		cpuInfo := string(cpuInfoBytes)
		cpuInfo = findColumnValue(cpuInfo, "model name")

		if cpuInfo != "" {
			return cpuInfo
		}
	}

	// Via "lscpu" (Linux)
	cpuInfoBytes, err = exec.Command("lscpu").Output()

	if err == nil {
		cpuInfo := string(cpuInfoBytes)
		cpuInfo = findColumnValue(cpuInfo, "Model name")

		if cpuInfo != "" {
			return cpuInfo
		}
	}

	// Via "sysctl" (Mac)
	cpuInfoBytes, err = exec.Command("sysctl", "-n", "machdep.cpu.brand_string").Output()

	if err == nil {
		cpuInfoBytes = bytes.TrimSpace(cpuInfoBytes)
		return string(cpuInfoBytes)
	}

	return ""
}

// findColumnValue finds the value of the given column inside a table.
func findColumnValue(output string, column string) string {
	columnPos := strings.Index(output, column)

	if columnPos == -1 {
		return ""
	}

	output = output[columnPos+len(column):]
	output = strings.TrimSpace(output)
	output = strings.TrimPrefix(output, ":")
	output = strings.TrimSpace(output)

	newline := strings.Index(output, "\n")

	if newline != -1 {
		output = output[:newline]
	}

	return output
}
