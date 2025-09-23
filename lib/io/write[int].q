write(n int) {
	if n < 0 {
		write("-")
	}

	writeInt(n, 10)
}

writeInt(n int, base int) {
	if n >= base || n <= -base {
		writeInt(n / base, base)
	}

	writeDigit(n % base)
}

writeDigit(n int) {
	write("FEDCBA9876543210123456789ABCDEF"[n+15..n+16])
}