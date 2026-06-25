write(n int) {
	if n < 0 {
		write("-")
	}

	writeDigits(n, 10)
}

write(n uint) {
	writeDigits(n, 10)
}

write(n *any) {
	writeDigits(n as uint, 16)
}