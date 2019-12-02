import sys

main() {
	fileName = "test.txt"
	contents = "123456789\n"
	length = 10

	file = sys.open(fileName, 66, 438)
	bytesWritten = sys.write(file, contents, length)
	sys.close(file)
	sys.unlink(fileName)

	if bytesWritten != length {
		sys.exit(1)
	}
}
