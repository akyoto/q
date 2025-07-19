import os

main() {
	x := 0

	if x != x && x != x {
		os.exit(1)
	}

	if x == x && x != x {
		os.exit(1)
	}

	if x != x && x == x {
		os.exit(1)
	}

	if x == x && x != x && x != x {
		os.exit(1)
	}

	if x != x && x == x && x != x {
		os.exit(1)
	}

	if x != x && x != x && x == x {
		os.exit(1)
	}

	if x == x && x == x && x == x {
		os.exit(0)
	}

	os.exit(1)
}