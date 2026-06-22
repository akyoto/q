main() {
	x := acquire()[1..]
}

acquire() -> ![]int {
	return new(int, 2)
}