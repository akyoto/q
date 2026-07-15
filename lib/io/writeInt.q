import strings

const {
	SIZE_SIGNED_10 = 20
	SIZE_UNSIGNED_10 = 19
	SIZE_UNSIGNED_16 = 16
}

write(n int) {
	buffer := new(byte, SIZE_SIGNED_10)
	num := strings.fromInt(n, 10, buffer)
	write(num)
}

write(n uint) {
	buffer := new(byte, SIZE_UNSIGNED_10)
	num := strings.fromInt(n, 10, buffer)
	write(num)
}

write(n *any) {
	buffer := new(byte, SIZE_UNSIGNED_16)
	num := strings.fromInt(n as uint, 16, buffer)
	write(num)
}