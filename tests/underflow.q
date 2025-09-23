import math

main() {
	x := math.minInt64
	assert x == math.minInt64
	x -= 1
	assert x == math.maxInt64
}