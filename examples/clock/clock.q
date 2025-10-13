import io
import time

main() {
	start := time.now()

	loop {
		elapsed := time.since(start)
		io.writeLine(elapsed / time.second)
		time.sleep(time.second - elapsed % time.second)
	}
}