factorial(n) {
	require n >= 0
	ensure _ >= 0

	mut x = 1

	for i = 1..n+1 {
		x = x * i
	}

	return x
}
