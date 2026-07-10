import io
import time

main() {
	loop 0..4 {
		go work()
	}

	time.sleep(time.second)
}

work() {
	io.write("[ ] start\n")
	time.sleep(500 * time.millisecond)
	io.write("[x] end\n")
}