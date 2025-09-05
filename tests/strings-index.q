import strings

main() {
	posEmpty, err := strings.index("Hello", "")
	assert err == 0
	assert posEmpty == 0

	posH, err := strings.index("Hello", "H")
	assert err == 0
	assert posH == 0

	posE, err := strings.index("Hello", "e")
	assert err == 0
	assert posE == 1

	posL, err := strings.index("Hello", "l")
	assert err == 0
	assert posL == 2

	posO, err := strings.index("Hello", "o")
	assert err == 0
	assert posO == 4

	posLL, err := strings.index("Hello", "ll")
	assert err == 0
	assert posLL == 2

	_, err := strings.index("Hello", "Hello World")
	assert err != 0

	_, err := strings.index("Hello", "hello")
	assert err != 0

	_, err := strings.index("Hello", "ella")
	assert err != 0
}