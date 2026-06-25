writeDigit(n int) {
	write("FEDCBA9876543210123456789ABCDEF"[n+15..n+16])
}

writeDigit(n uint) {
	write("0123456789ABCDEF"[n..n+1])
}