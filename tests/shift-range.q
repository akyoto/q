main() {
	assert 1 << 65 == 2
	assert 1 << 592 == 1 << 16
	assert 1 << -1 == (-(9223372036854775807) - 1)
	assert 1 >> -7 == 0
	assert 76 >> 592 == 0
	assert 76 >> 64 == 76
	assert 123 >> 64 == 123
}