import mem
import strings

run(path string) -> error {
	pid := fork()

	if pid == 0 {
		cpath := strings.c(path)
		err := execve(cpath.ptr, 0, 0)
		mem.free(cpath)
		return err
	}

	return waitid(1, pid, 0, 0x4)
}

execve(path *byte, argv *any, envp *any) -> error {
	return syscall(_execve, path, argv, envp)
}

waitid(type int, id int, info *any, options int) -> error {
	return syscall(_waitid, type, id, info, options)
}