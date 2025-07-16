main() {
	f1(1, 2, 3, 4, 5, 6)
}

f1(a int, b int, c int, d int, e int, f int) {
	f2(f, e, d, c, b, a)
}

f2(a int, b int, c int, d int, e int, f int) {
	assert a == 6
	assert b == 5
	assert c == 4
	assert d == 3
	assert e == 2
	assert f == 1
}