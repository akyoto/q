main() {
	a := 1
	b := 2
	c := a < b
	d := b < a
	e := a > 0
	checkFalse(d)
	checkTrue(c)
	checkTrue(e)
}

checkFalse(v bool) {
	assert !v
}

checkTrue(v bool) {
	assert v
}