package verbose_test

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/verbose"
	"git.urbach.dev/go/assert"
	"git.urbach.dev/go/color"
)

func TestVerboseOutput(t *testing.T) {
	examples := "../../examples"

	fs.Walk(examples, func(dir string) {
		t.Run(dir, func(t *testing.T) {
			b := config.New(filepath.Join(examples, dir))
			test(t, b)
		})
	})

	tests := "../../tests"

	fs.Walk(tests, func(file string) {
		if !strings.HasSuffix(file, ".q") {
			return
		}

		t.Run(file, func(t *testing.T) {
			b := config.New(filepath.Join(tests, file))
			b.Lint(false)
			test(t, b)
		})
	})
}

func test(t *testing.T, b *config.Build) {
	originalOut := ""

	for range 3 {
		env, err := compiler.Compile(b)
		assert.Nil(t, err)

		// Capture output
		stdout := os.Stdout
		reader, writer, err := os.Pipe()
		assert.Nil(t, err)
		color.Redirect(writer)
		channel := make(chan string)

		go func() {
			out, err := io.ReadAll(reader)
			assert.Nil(t, err)
			assert.NotNil(t, out)
			channel <- string(out)
		}()

		verbose.ASM(env)
		verbose.SSA(env)

		writer.Close()
		color.Redirect(stdout)
		out := <-channel

		if originalOut == "" {
			originalOut = out
		} else {
			assert.Equal(t, out, originalOut)
		}
	}
}