main() {
	n := 10
	x := 1

	loop {
		if n == 0 {
			return
		}

		f(x)
		n = n - 1
	}
}

f(x int) -> int {
	return x
}