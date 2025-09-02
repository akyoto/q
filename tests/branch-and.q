import run

main() {
	x := 0

	if x != x && x != x {
		run.exit(1)
	}

	if x == x && x != x {
		run.exit(1)
	}

	if x != x && x == x {
		run.exit(1)
	}

	if x == x && x != x && x != x {
		run.exit(1)
	}

	if x != x && x == x && x != x {
		run.exit(1)
	}

	if x != x && x != x && x == x {
		run.exit(1)
	}

	if x == x && x == x && x == x {
		run.exit(0)
	}

	run.exit(1)
}