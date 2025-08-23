write(n int) {
	if n < 0 {
		write("-")
		n = -n
	}

	if n >= 10 {
		write(n / 10)
	}

	writeDigit(n % 10)
}
