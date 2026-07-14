import c
import strings

env(name string) -> (value string, err error) {
	ptr := envp

	loop {
		if [ptr] == 0 {
			return "", -1
		}

		current := string{ptr: ptr, len: c.length(ptr)}
		value, err := extract(current, name)

		if err == 0 {
			return value, 0
		}

		ptr += current.len + 1
	}
}

extract(current string, name string) -> (string, error) {
	key, value, err := strings.cut(current, "=")
	assert err == 0

	if key == name {
		return value, 0
	}

	return "", -1
}

global {
	envp *byte
}