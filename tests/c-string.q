import c

main() {
	s := c.string("")
	assert s.len == 1
	assert s[s.len-1] == 0
	delete(s)

	s := c.string("H")
	assert s.len == 2
	assert s[s.len-1] == 0
	delete(s)

	s := c.string("He")
	assert s.len == 3
	assert s[s.len-1] == 0
	delete(s)

	s := c.string("Hel")
	assert s.len == 4
	assert s[s.len-1] == 0
	delete(s)

	s := c.string("Hell")
	assert s.len == 5
	assert s[s.len-1] == 0
	delete(s)

	s := c.string("Hello")
	assert s.len == 6
	assert s[s.len-1] == 0
	delete(s)
}