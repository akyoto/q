import {
	"sys"
}

main() {
	fileName = "test.txt"
	contents = "123456789\n"
	length = 10

	file = sys.open(fileName)
	bytesWritten = sys.write(file, contents, length)
	sys.close(file)
	sys.unlink(fileName)

	sys.write(1, contents, length)

	if bytesWritten == length {
		sys.exit(0)
	}

	sys.exit(1)
}
