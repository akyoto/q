package data_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/data"
	"git.urbach.dev/go/assert"
)

func TestInterning(t *testing.T) {
	d := data.Data{}
	d.SetImmutable("label1", []byte("Hello"))
	d.SetImmutable("label2", []byte("ello"))
	raw, positions := d.Finalize()
	assert.DeepEqual(t, raw, []byte("Hello"))
	assert.Equal(t, positions["label1"], 0)
	assert.Equal(t, positions["label2"], 1)
}

func TestInterningReverse(t *testing.T) {
	d := data.Data{}
	d.SetImmutable("label1", []byte("ello"))
	d.SetImmutable("label2", []byte("Hello"))
	raw, positions := d.Finalize()
	assert.DeepEqual(t, raw, []byte("Hello"))
	assert.Equal(t, positions["label1"], 1)
	assert.Equal(t, positions["label2"], 0)
}

func TestNoInterning(t *testing.T) {
	d := data.Data{}
	d.SetMutable("label1", []byte{0xAB, 0xCD, 0xEF})
	d.SetMutable("label2", []byte{0xCD, 0xEF})
	raw, positions := d.Finalize()
	assert.DeepEqual(t, raw, []byte{0xAB, 0xCD, 0xEF, 0x00, 0xCD, 0xEF})
	assert.Equal(t, positions["label1"], 0)
	assert.Equal(t, positions["label2"], 4)
}