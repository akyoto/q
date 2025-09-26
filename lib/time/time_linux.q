now() -> int {
	t := new(Timespec)
	syscall(_clock_gettime, realtime, t)
	n := t.seconds * second + t.nanoseconds
	delete(t)
	return n
}

sleep(nanoseconds int) {
	seconds := 0

	if nanoseconds >= second {
		seconds = nanoseconds / second
		nanoseconds = nanoseconds % second
	}

	duration := new(Timespec)
	duration.seconds = seconds
	duration.nanoseconds = nanoseconds
	syscall(_nanosleep, duration)
	delete(duration)
}