import c
import io
import mem

writeFile(path string, content string) -> error {
	cpath := c.string(path)
	fd, err := openWrite(cpath.ptr)
	mem.free(cpath)

	if err != 0 {
		return err
	}

	io.writeTo(fd, content)
	close(fd)
	return 0
}