Mixed {
	a uint8
	b int32
	c int16
}

main() {
	f(Mixed{a: 0, b: -1, c: 0x1234})
}

f(m Mixed) {
	assert m.a == 0
	assert m.b == -1
	assert m.c == 0x1234
}