import "sys"

main() {
	fileName = "test.txt"
	contents = "123456789\n"
	length = 10

	file = open(fileName)
	bytesWritten = write(file, contents, length)
	close(file)
	unlink(fileName)

	if bytesWritten != length {
		exit(1)
	}

	exit(0)
}
