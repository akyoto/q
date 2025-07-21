main() {
	assert 0b00 & 0b00 == 0b00
	assert 0b00 & 0b01 == 0b00
	assert 0b01 & 0b00 == 0b00
	assert 0b01 & 0b01 == 0b01
	assert 0b01 & 0b10 == 0b00
	assert 0b01 & 0b11 == 0b01
	assert 0b10 & 0b00 == 0b00
	assert 0b10 & 0b01 == 0b00
	assert 0b10 & 0b10 == 0b10
	assert 0b10 & 0b11 == 0b10
	assert 0b11 & 0b00 == 0b00
	assert 0b11 & 0b01 == 0b01
	assert 0b11 & 0b10 == 0b10
	assert 0b11 & 0b11 == 0b11
}