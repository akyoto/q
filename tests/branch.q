import run

main() {
	x := 0

	if x != 0 {
		run.exit(1)
	}

	if x > 0 {
		run.exit(1)
	}

	if x < 0 {
		run.exit(1)
	}

	if 0 != x {
		run.exit(1)
	}

	if 0 > x {
		run.exit(1)
	}

	if 0 < x {
		run.exit(1)
	}

	if x >= 1 {
		run.exit(1)
	}

	if 1 <= x {
		run.exit(1)
	}

	if x + 1 != x + 1 {
		run.exit(1)
	}

	if x + 1 != inc(x) {
		run.exit(1)
	}

	if x - 1 != dec(x) {
		run.exit(1)
	}

	if inc(x) != x + 1 {
		run.exit(1)
	}

	if dec(x) != x - 1 {
		run.exit(1)
	}

	if x != inc(dec(x)) {
		run.exit(1)
	}

	if inc(dec(x)) != x {
		run.exit(1)
	}

	if inc(x) == dec(x) {
		run.exit(1)
	}

	if x == 0 {
		run.exit(0)
	}

	run.exit(1)
}

inc(x int) -> int {
	return x + 1
}

dec(x int) -> int {
	return x - 1
}