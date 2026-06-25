writeDigits(n int, base int) {
	if n >= base || n <= -base {
		writeDigits(n / base, base)
	}

	writeDigit(n % base)
}

writeDigits(n uint, base uint) {
	if n >= base {
		writeDigits(n / base, base)
	}

	writeDigit(n % base)
}