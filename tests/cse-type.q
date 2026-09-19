main() {
	a := -100
	b := 5 as uint
	c := a + b
	d := b + a
	assert c >> 60 == -1
	assert d >> 60 == 15
}