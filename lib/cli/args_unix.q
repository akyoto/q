import c

args() -> []string {
	count := stack[0]
	args := new(string, count)

	loop i := 0..count {
		ptr := stack[1+i] as *byte
		args[i] = string{ptr: ptr, len: c.length(ptr)}
	}

	return args
}