main() {
	assert negative(1) == -1
	assert negative(-1) == 1
	assert negative(256) == -256
	assert negative(-256) == 256
	assert negative(1152921504606846976) == -1152921504606846976
	assert negative(-1152921504606846976) == 1152921504606846976
}

negative(x int) -> int {
	return -x
}