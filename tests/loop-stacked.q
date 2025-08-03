import io
import os

main() {
	n := 10

	loop {
		loop {
			if n == 0 {
				return
			}

			io.write(".")
			n -= 1
		}
	}

	os.exit(1)
}