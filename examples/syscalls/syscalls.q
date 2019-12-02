import fs
import process

main() {
	fileName = "test.txt"
	contents = "123456789\n"
	length = 10

	file = fs.open(fileName)
	bytesWritten = fs.write(file, contents, length)
	fs.close(file)
	fs.unlink(fileName)

	if bytesWritten != length {
		process.exit(1)
	}
}
