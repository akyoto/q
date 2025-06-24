package data_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/data"
	"git.urbach.dev/go/assert"
)

func TestInterning(t *testing.T) {
	d := data.Data{}
	d.Insert("label1", []byte("Hello"))
	d.Insert("label2", []byte("ello"))
	raw, positions := d.Finalize()
	assert.DeepEqual(t, raw, []byte("Hello"))
	assert.Equal(t, positions["label1"], 0)
	assert.Equal(t, positions["label2"], 1)
}

func TestInterningReverse(t *testing.T) {
	d := data.Data{}
	d.Insert("label1", []byte("ello"))
	d.Insert("label2", []byte("Hello"))
	raw, positions := d.Finalize()
	assert.DeepEqual(t, raw, []byte("Hello"))
	assert.Equal(t, positions["label1"], 1)
	assert.Equal(t, positions["label2"], 0)
}