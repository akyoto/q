import sys

deleteFile(fileName Text) {
	#require fileName != ""
	sys.unlink(fileName)
}
