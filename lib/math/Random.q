Random {
	s0 uint64
	s1 uint64
}

next(rng *Random) -> uint64 {
	s0 := rng.s0
	s1 := rng.s1
	result := rotateLeft(s0 + s1, 17) + s0
	s1 ^= s0
	rng.s0 = rotateLeft(s0, 49) ^ s1 ^ (s1 << 21)
	rng.s1 = rotateLeft(s1, 28)
	return result
}