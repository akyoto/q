import mem
import net

address(path string) -> !string {
	addr := mem.alloc(2 + path.len + 1)
	addr[0] = addr.len as byte
	addr[1] = net.AF_UNIX
	mem.copy(addr[2..addr.len-1], path)
	return addr
}