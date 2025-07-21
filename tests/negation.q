main() {
	assert negative(1) == -1
	assert negative(-1) == 1
	assert negative(256) == -256
	assert negative(-256) == 256
}

negative(x int) -> int {
	return -x
}