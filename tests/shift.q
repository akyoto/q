main() {
	assert 0 << 0 == 0
	assert 0 >> 0 == 0
	assert 1 << 0 == 1
	assert 1 >> 0 == 1
	assert 1 >> 1 == 0
	assert 1 << 1 == 2
	assert 1 << 2 == 4
	assert 1 << 3 == 8
	assert 1 << 4 == 16
	assert 16 >> 1 == 8
	assert 16 >> 2 == 4
	assert 16 >> 3 == 2
	assert 16 >> 4 == 1
	assert -16 >> 1 == -8
	assert -16 >> 2 == -4
	assert -16 >> 3 == -2
	assert -16 >> 4 == -1
}