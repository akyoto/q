package tests_test

import (
	"hash/crc32"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/cli/q/src/linker"
	"git.urbach.dev/go/assert"
)

type run struct {
	Name     string
	Args     []string
	Input    string
	Output   string
	ExitCode int
}

// Run compiles a debug and release version and tests both.
func (test *run) Run(t *testing.T, build *config.Build) {
	build.Matrix(func(cross *config.Build) {
		if cross.OS == build.OS && cross.Arch == build.Arch {
			cross.Optimize(false)
			test.RunBuild(t, test.Name+"/"+cross.OS.String()+"/"+cross.Arch.String()+"/debug", cross)
			cross.Optimize(true)
			test.RunBuild(t, test.Name+"/"+cross.OS.String()+"/"+cross.Arch.String()+"/release", cross)
		} else {
			cross.Optimize(false)
			test.Compile(t, test.Name+"/"+cross.OS.String()+"/"+cross.Arch.String()+"/debug", cross)
			cross.Optimize(true)
			test.Compile(t, test.Name+"/"+cross.OS.String()+"/"+cross.Arch.String()+"/release", cross)
		}
	})
}

// Compile only tests the compilation without actually running the executable.
func (test *run) Compile(t *testing.T, name string, build *config.Build) {
	t.Run(name, func(t *testing.T) {
		env, err := compiler.Compile(build)
		assert.Nil(t, err)
		discard := &exe.Discard{}
		linker.Write(discard, env)
	})
}

// RunBuild builds and runs the file to check if the output matches the expected output.
func (test *run) RunBuild(t *testing.T, name string, build *config.Build) {
	t.Run(name, func(t *testing.T) {
		if test.ExitCode == -1 {
			env, err := compiler.Compile(build)
			assert.Nil(t, err)
			discard := &exe.Discard{}
			linker.Write(discard, env)
			return
		}

		originalHash := uint32(0)
		tmpDir := os.TempDir()
		err := os.MkdirAll(tmpDir, 0o755)
		assert.Nil(t, err)
		executable := build.Executable()
		executable = filepath.Join(tmpDir, filepath.Base(executable))

		for range 3 {
			env, err := compiler.Compile(build)
			assert.Nil(t, err)

			err = linker.WriteFile(executable, env)
			assert.Nil(t, err)

			stat, err := os.Stat(executable)
			assert.Nil(t, err)
			assert.True(t, stat.Size() > 0)

			// Run the executable
			cmd := exec.Command(executable, test.Args...)
			cmd.Stdin = strings.NewReader(test.Input)
			output, err := cmd.Output()
			exitCode := 0

			if err != nil {
				exitError, ok := err.(*exec.ExitError)

				if !ok {
					t.Fatal(exitError)
				}

				exitCode = exitError.ExitCode()
			}

			assert.Equal(t, exitCode, test.ExitCode)
			assert.DeepEqual(t, string(output), test.Output)

			// Fail the test if the machine code is not deterministic
			newHash, err := checksum(executable)
			assert.Nil(t, err)

			if originalHash == 0 {
				originalHash = newHash
			} else {
				assert.Equal(t, newHash, originalHash)
			}

			// Clean up
			err = os.Remove(executable)
			assert.Nil(t, err)
		}
	})
}

// checksum calculates a checksum for the file contents.
func checksum(path string) (uint32, error) {
	file, err := os.Open(path)

	if err != nil {
		return 0, err
	}

	defer file.Close()
	sum := crc32.NewIEEE()
	_, err = io.Copy(sum, file)

	if err != nil {
		return 0, err
	}

	return sum.Sum32(), nil
}