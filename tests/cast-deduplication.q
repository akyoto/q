main() {
	x := -1
	a := x as uint64
	b := x as int32
	assert a >> 31 == 0x1FFFFFFFF
	assert b >> 31 == -1
}