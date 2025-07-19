main() {
	assert '\0' == 0
	assert '\t' == 9
	assert '\n' == 10
	assert '\r' == 13
	assert '\"' == 34
	assert '\'' == 39
	assert '\\' == 92
	assert '0' == 48
	assert 'A' == 65
	assert 'a' == 97
	assert '世' == 0x4E16
	assert '界' == 0x754C
	assert '😀' == 0x1F600
}