import sys

writeFile(fileName Text, contents Text, length Int) {
	#require fileName != ""
	file = sys.open(fileName, 66, 438)
	sys.write(file, contents, length)
	sys.close(file)
}
