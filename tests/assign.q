main() {
	a := 10
	b := 0
	assert a != b
	b = a
	assert a == b
	a -= 2
	assert a < b
	a += 2
	assert a == b
	a *= 2
	assert a > b
	a /= 2
	assert a == b
	a <<= 2
	assert a > b
	a >>= 2
	assert a == b
	a |= 1
	assert a == 11
	a &= 1
	assert a == 1
	a ^= 1
	assert a == 0
	a %= 2
	assert a < 2
}