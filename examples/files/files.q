import fs

main() {
	fileName := "test.txt"
	contents := "123456789\n"
	length := 10

	fs.writeFile(fileName, contents, length)
	fs.deleteFile(fileName)
}
