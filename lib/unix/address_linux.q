import mem

address(path string) -> !string {
	addr := mem.alloc(2 + path.len + 1)
	addr[0] = AF_UNIX
	mem.copy(addr[2..addr.len-1], path)
	return addr
}