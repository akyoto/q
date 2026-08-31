Mixed {
	a int8
	b int32
}

main() {
	f(Mixed{a: -1, b: 0x12345678})
}

f(m Mixed) {
	assert m.a == -1
	assert m.b == 0x12345678
}