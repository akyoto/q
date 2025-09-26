import io
import time

main() {
	start := time.now()

	loop {
		elapsed := time.since(start)
		io.write(elapsed / time.second)
		io.write("\n")
		time.sleep(time.second - elapsed % time.second)
	}
}