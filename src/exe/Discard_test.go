package exe_test

import (
	"io"
	"testing"

	"git.urbach.dev/cli/q/src/exe"
)

func TestDiscard(t *testing.T) {
	discard := exe.Discard{}
	discard.Write(nil)
	discard.Seek(0, io.SeekCurrent)
}