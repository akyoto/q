import os

write(buffer []byte) -> (written int) {
	return os.write(0, buffer, len(buffer))
}