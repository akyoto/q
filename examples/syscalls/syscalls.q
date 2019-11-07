main() {
	write("Hello Syscalls", 14)
}

write(msg, length) {
	syscall(1, 1, msg, length)
}
