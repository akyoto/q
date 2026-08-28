main() {
	total := 0

	loop i := 0..4 {
		total += i
	}

	total += f(3)
}

f(x int) -> int {
	return x * 2
}