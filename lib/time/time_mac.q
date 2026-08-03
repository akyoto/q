now() -> int {
	t := new(Timeval)
	syscall(_gettimeofday, t, 0, 0)
	n := t.seconds * second + t.microseconds * microsecond
	return n
}

sleep(nanoseconds int) {
	seconds := 0

	if nanoseconds >= second {
		seconds = nanoseconds / second
		nanoseconds = nanoseconds % second
	}

	duration := new(Timespec) {
		seconds: seconds,
		nanoseconds: nanoseconds,
	}

	System.nanosleep(duration, 0)
}

extern {
	System {
		nanosleep(requested *Timespec, remaining *Timespec|nil) -> int
	}
}