import c

args() -> []string {
	args := new(string, argc)

	loop i := 0..argc {
		ptr := argv[i]
		args[i] = string{ptr: ptr, len: c.length(ptr)}
	}

	return args
}

global {
	argc uint
	argv **byte
}