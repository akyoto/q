main() {
	a := -1
	b := -2

	assert a == -1
	assert a != 0xFF
	assert a != 0xFFFF
	assert a != 0xFFFFFFFF
	assert b == -2
	assert b != 0xFE
	assert b != 0xFFFE
	assert b != 0xFFFFFFFE
	assert a + b == -3
	assert a - b == 1
	assert a * b == 2
	assert a / b == 0
	assert a % b == -1
}