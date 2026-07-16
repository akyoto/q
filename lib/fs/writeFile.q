import c
import io

writeFile(path string, buffer string) -> error {
	cpath := c.string(path)
	fd, err := openWrite(cpath.ptr)
	delete(cpath)

	if err != 0 {
		return err
	}

	pos := 0

	loop {
		n, err := io.writeTo(fd, buffer[pos..])

		if err != 0 {
			close(fd)
			return err
		}

		pos += n

		if pos >= buffer.len {
			close(fd)
			return 0
		}
	}
}