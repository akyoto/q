writeLine(buffer string) {
	write(buffer)
	write("\n")
}

writeLine(signed int) {
	write(signed)
	write("\n")
}

writeLine(unsigned uint) {
	write(unsigned)
	write("\n")
}

writeLine(pointer *any) {
	write(pointer)
	write("\n")
}