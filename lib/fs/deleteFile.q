import sys

deleteFile(fileName Text) {
	#expect fileName != ""
	sys.unlink(fileName)
}
