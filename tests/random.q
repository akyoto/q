import math

main() {
	rand := new(math.Random) {
		s0: 0x2BD7A6A6E99C2DDC,
		s1: 0x0992CCAF6A6FCA05
	}

	assert math.next(rand) == 0x12844EBED95E98B0
	assert math.next(rand) == 0x3D7CD35F4CA1D6FA
	assert math.next(rand) == 0xC0EFFFEBCADDF9FE

	delete(rand)
}