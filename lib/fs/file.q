import sys

writeFile(fileName, contents, length) {
	file = sys.open(fileName, 66, 438)
	sys.write(file, contents, length)
	sys.close(file)
}

deleteFile(fileName) {
	sys.unlink(fileName)
}
