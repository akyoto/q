main() {
	x := new(int, 8)

	x[0] = 10
	x[0] += 5
	assert x[0] == 15

	x[1] = 10
	x[1] -= 4
	assert x[1] == 6

	x[2] = 3
	x[2] *= 4
	assert x[2] == 12

	x[3] = 20
	x[3] /= 3
	assert x[3] == 6

	x[4] = 17
	x[4] %= 5
	assert x[4] == 2

	x[5] = 0xFF
	x[5] &= 0x0F
	assert x[5] == 0x0F

	x[6] = 0xF0
	x[6] |= 0x0F
	assert x[6] == 0xFF

	x[7] = 0xF0
	x[7] ^= 0x0F
	assert x[7] == 0xFF
}