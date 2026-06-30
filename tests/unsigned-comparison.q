main() {
	a := 0x80 as uint8
	assert a > 0
	b := 0x8000 as uint16
	assert b > 0
	c := 0x80000000 as uint32
	assert c > 0
	d := 0x8000000000000000 as uint64
	assert d > 0
}