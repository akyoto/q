import bits

Random {
	s0 uint64
	s1 uint64
}

newRandom(seed uint64) -> !*Random {
	rng := new(Random)
	s0, next := splitMix64(seed)
	rng.s0 = s0
	s1, _ := splitMix64(next)
	rng.s1 = s1
	return rng
}

next(rng *Random) -> uint64 {
	s0 := rng.s0
	s1 := rng.s1
	result := bits.rotateLeft(s0 + s1, 17) + s0
	s1 ^= s0
	rng.s0 = bits.rotateLeft(s0, 49) ^ s1 ^ (s1 << 21)
	rng.s1 = bits.rotateLeft(s1, 28)
	return result
}

splitMix64(state uint64) -> (result uint64, next uint64) {
	state += 0x9E3779B97F4A7C15
	z := state
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	return z ^ (z >> 31), state
}