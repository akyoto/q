import sys

writeFile(fileName, contents, length) {
	#require fileName != ""
	file = sys.open(fileName, 66, 438)
	sys.write(file, contents, length)
	sys.close(file)
}
