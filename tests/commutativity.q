main() {
	loop a := -10..10 {
		loop b := -10..10 {
			x := a + b
			y := b + a
			assert x == y
			x = a * b
			y = b * a
			assert x == y
			x = a & b
			y = b & a
			assert x == y
			x = a | b
			y = b | a
			assert x == y
			x = a ^ b
			y = b ^ a
			assert x == y
		}
	}
}