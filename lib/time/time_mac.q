import io

now() -> int {
	t := new(Timespec)
	syscall(_gettimeofday, t, 0, 0)
	n := t.seconds * second + t.microseconds * microsecond
	delete(t)
	return n
}

sleep(_nanoseconds int) {
	io.write("time.sleep: not implemented on Mac\n")
}