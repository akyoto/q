main() {
	fileName = "test.txt"
	contents = "123456789\n"
	length = 10

	file = open(fileName)
	bytesWritten = write(file, contents, length)
	close(file)
	unlink(fileName)

	write(1, contents, length)

	if bytesWritten == length {
		exit(0)
	}

	exit(1)
}

open(fileName) {
	return syscall(2, fileName, 66, 438)
}

unlink(fileName) {
	return syscall(87, fileName)
}

write(fd, msg, length) {
	return syscall(1, fd, msg, length)
}

close(fd) {
	return syscall(3, fd)
}

exit(code) {
	syscall(60, code)
}
