import c

main() {
	assert c.length("\0".ptr) == 0
	assert c.length("H\0".ptr) == 1
	assert c.length("He\0".ptr) == 2
	assert c.length("Hel\0".ptr) == 3
	assert c.length("Hell\0".ptr) == 4
	assert c.length("Hello\0".ptr) == 5
}