import sys

deleteFile(fileName) {
	#require fileName != ""
	sys.unlink(fileName)
}
