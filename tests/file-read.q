import fs
import mem

main() {
	source, err := fs.readFile("file-read.q")
	assert err == 0
	assert source.len > 150
	assert source.len < 200
	mem.free(source)
}