import io
import mem
import strings

readFile(path string) -> (!string, error) {
	cpath := strings.c(path)
	fd, err := open(cpath.ptr, 0, 0)
	mem.free(cpath)

	if err != 0 {
		return "", err
	}

	fileSize, err := size(fd)

	if err != 0 {
		close(fd)
		return "", err
	}

	buffer := mem.alloc(fileSize)
	pos := 0

	loop {
		n := io.readFrom(fd, buffer[pos..])

		if n <= 0 {
			close(fd)
			return buffer, 0
		}

		pos += n
	}
}