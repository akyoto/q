global {
	data [10]int
}

main() {
	take(data)
}

take(a []int) {
	assert a.len == 10
}