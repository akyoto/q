main() {
	loop i := 0..10 {
		assert fac1(i) == fac2(i)
	}
}

fac1(n int) -> int {
	x := 1

	loop i := 1..n+1 {
		x *= i
	}

	return x
}

fac2(n int) -> int {
	if n <= 1 {
		return 1
	}

	return n * fac2(n - 1)
}