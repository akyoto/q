main() {
	assert left(10) == 30
	assert right(10) == 30
	assert leftSwap(10) == 30
	assert rightSwap(10) == 30
}

left(x int) -> int {
	return (x + 10) + 10
}

leftSwap(x int) -> int {
	return (10 + x) + 10
}

right(x int) -> int {
	return 10 + (10 + x)
}

rightSwap(x int) -> int {
	return 10 + (x + 10)
}