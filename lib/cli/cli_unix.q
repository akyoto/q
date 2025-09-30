import c
import run

args() -> []string {
	count := run.stack[0]
	args := new(string, count)

	loop i := 0..count {
		ptr := run.stack[1+i] as *byte
		args[i] = string{ptr: ptr, len: c.length(ptr)}
	}

	return args
}