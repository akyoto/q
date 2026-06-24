main() {
	x = 10
	assert x == 10
	x += 1
	assert x == 11
	x -= 1
	assert x == 10
	x *= 2
	assert x == 20
	x /= 2
	assert x == 10
}

global {
	x int
}