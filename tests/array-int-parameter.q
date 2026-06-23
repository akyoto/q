main() {
	x := new(int, 10)
	x[0] = 42
	assert first(x) == 42
}

first(x []int) -> int {
    return x[0]
}