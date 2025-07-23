import os

main() {
	n := 10

	loop {
		if n == 0 {
			os.exit(0)
		}

		n = n - 1
	}

	os.exit(1)
}