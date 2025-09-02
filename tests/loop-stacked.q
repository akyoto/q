import io
import run

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

	run.exit(1)
}