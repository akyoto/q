local {
	buffer [20]byte
}

main() {
	buffer[3] = 3
	assert buffer[3] == 3
	x := new(byte, 20000000)
	assert buffer[3] == 3
}