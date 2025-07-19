import os

main() {
	x := 0

	if x != 0 {
		os.exit(1)
	}

	if x > 0 {
		os.exit(1)
	}

	if x < 0 {
		os.exit(1)
	}

	if 0 != x {
		os.exit(1)
	}

	if 0 > x {
		os.exit(1)
	}

	if 0 < x {
		os.exit(1)
	}

	if x >= 1 {
		os.exit(1)
	}

	if 1 <= x {
		os.exit(1)
	}

	if x + 1 != x + 1 {
		os.exit(1)
	}

	if x + 1 != inc(x) {
		os.exit(1)
	}

	if x - 1 != dec(x) {
		os.exit(1)
	}

	if inc(x) != x + 1 {
		os.exit(1)
	}

	if dec(x) != x - 1 {
		os.exit(1)
	}

	if x != inc(dec(x)) {
		os.exit(1)
	}

	if inc(dec(x)) != x {
		os.exit(1)
	}

	if inc(x) == dec(x) {
		os.exit(1)
	}

	if x == 0 {
		os.exit(0)
	}

	os.exit(1)
}

inc(x int) -> int {
	return x + 1
}

dec(x int) -> int {
	return x - 1
}