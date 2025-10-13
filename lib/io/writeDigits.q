write(n int) {
	if n < 0 {
		write("-")
	}

	writeDigits(n, 10)
}

writeDigits(n int, base int) {
	if n >= base || n <= -base {
		writeDigits(n / base, base)
	}

	writeDigit(n % base)
}

writeDigit(n int) {
	write("FEDCBA9876543210123456789ABCDEF"[n+15..n+16])
}