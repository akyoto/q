import c

args() -> []string {
	count := argc
	pointers := argv
	args := new(string, count)

	loop i := 0..count {
		ptr := pointers[i]
		args[i] = string{ptr: ptr, len: c.length(ptr)}
	}

	return args
}

global {
	argc uint
	argv **byte
}