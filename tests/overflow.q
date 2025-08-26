import math

main() {
	x := math.maxInt64
	assert x == math.maxInt64
	x += 1
	assert x < 0
}