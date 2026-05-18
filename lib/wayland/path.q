import cli
import mem

path() -> (!string, error) {
	runtime, err := cli.env("XDG_RUNTIME_DIR")

	if err != 0 {
		return "", -1
	}

	display, err := cli.env("WAYLAND_DISPLAY")

	if err != 0 {
		return "", -1
	}

	address := mem.alloc(runtime.len + 1 + display.len)
	mem.copy(address, runtime)
	address[runtime.len] = '/'
	mem.copy(address[runtime.len+1..], display)
	return address, 0
}