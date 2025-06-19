import os

write(buffer []byte) -> int {
	return os.write(0, buffer, len(buffer))
}