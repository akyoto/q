main() {
	n := Number{value: 1}
	assert n.plus(0) == 1
	assert n.plus(1) == 2
	assert n.plus(2) == 3
	assert n.plus(3) == 4
}

Number {
	value int
}

plus(n Number, amount int) -> int {
	return n.value + amount
}