main() {
	loop i := 0..10 {
		assert fib1(i) == fib2(i)
	}
}

fib1(n int) -> int {
	b := 0
	c := 1

	loop 0..n {
		a := b
		b = c
		c = a + b
	}

	return b
}

fib2(n int) -> int {
	if n <= 1 {
		return n
	}

	return fib2(n - 1) + fib2(n - 2)
}