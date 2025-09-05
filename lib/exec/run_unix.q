import mem
import run
import strings

run(path string) -> error {
	cpath := strings.c(path)
	pid := fork()

	if pid == 0 {
		err := execve(cpath.ptr, 0, 0)
		run.exit(err)
	}

	status := new(int32)
	result := wait4(pid, status, 0, 0)

	if result < 0 {
		mem.free(cpath)
		return result
	}

	mem.free(cpath)
	return ([status] >> 8) & 0xFF
}

execve(path *byte, argv *any, envp *any) -> error {
	return syscall(_execve, path, argv, envp)
}

wait4(pid int, status *int32, options int, rusage *any) -> error {
	return syscall(_wait4, pid, status, options, rusage)
}