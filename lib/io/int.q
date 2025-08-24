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

writeDigit(n int) {
	write("0123456789"[n..n+1])
}