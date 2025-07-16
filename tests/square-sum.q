main() {
	assert f(2, 3) == 25
}

f(x int, y int) -> int {
	return (x + y) * (x + y)
}