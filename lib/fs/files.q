import sys

writeFile(fileName Text, contents Text, length Int) {
	#expect fileName != ""
	file := sys.open(fileName, 66, 438)
	sys.write(file, contents, length)
	sys.close(file)
}

deleteFile(fileName Text) {
	#expect fileName != ""
	sys.unlink(fileName)
}
