main() {
	f(1, 2)
}

f(x int, y int) {
	assert x << y == 0b100
}