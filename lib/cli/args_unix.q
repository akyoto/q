import c

args() -> ![]string {
	count := argc - 1
	pointers := argv
	args := new(string, count)

	loop i := 0..count {
		ptr := pointers[i+1]
		args[i] = string{ptr: ptr, len: c.length(ptr)}
	}

	return args
}

global {
	argc uint
	argv **byte
}