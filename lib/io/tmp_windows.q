read2(buffer *byte, length int) -> int {
	stdin := kernel32.GetStdHandle(-10)
	kernel32.ReadFile(stdin, buffer, length, 0, 0)
	return length
}

write2(buffer *byte, length int) -> (written int) {
	stdout := kernel32.GetStdHandle(-11)
	kernel32.WriteFile(stdout, buffer, length, 0, 0)
	return length
}