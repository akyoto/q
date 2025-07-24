import io
import os

main() {
	n := 5

	loop {
		if n == 0 {
			os.exit(0)
		}

		io.write(".")
		n = n - 1
	}

	os.exit(1)
}