import math

main() {
	rand := math.newRandom(0)
	assert math.next(rand) == 0x6F68E1E7E2646EE1
	assert math.next(rand) == 0xBF971B7F454094AD
	assert math.next(rand) == 0x48F2DE556F30DE38
	assert math.next(rand) == 0x6EA7C59F89BBFC75
	assert math.next(rand) == 0x765437C08F02E2F5
	delete(rand)
}