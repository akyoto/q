read2(buffer *byte, length int) -> int {
	stdin := kernel32.GetStdHandle(-10)
	count := new(uint32)
	kernel32.ReadFile(stdin, buffer, length, count, 0)
	return [count]
}

write2(buffer *byte, length int) -> (written int) {
	stdout := kernel32.GetStdHandle(-11)
	count := new(uint32)
	kernel32.WriteFile(stdout, buffer, length, count, 0)
	return [count]
}