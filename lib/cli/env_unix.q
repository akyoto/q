import c
import strings

env(name string) -> (value string, err error) {
	cur := envp

	loop {
		ptr := [cur]

		if ptr == 0 {
			return "", -1
		}

		current := string{ptr: ptr, len: c.length(ptr)}
		key, value, err := strings.cut(current, "=")
		assert err == 0

		if key == name {
			return value, 0
		}

		cur = cur + 8
	}
}

global {
	envp **byte
}