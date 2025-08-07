writeInt(n int) {
	if n < 0 {
		write("-")
		n = -n
	}

	if n >= 10 {
		writeInt(n / 10)
	}

	writeDigit(n % 10)
}