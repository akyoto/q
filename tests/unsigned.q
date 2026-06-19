main() {
	a := 0xFFFF as uint16
	assert a / 0x100 == 0xFF
	assert a % 0x100 == 0xFF
	assert a >> 8 == 0xFF

	b := 0x8000000000000000 as uint64
	assert b / 2 == 0x4000000000000000
	assert b % 2 == 0
}