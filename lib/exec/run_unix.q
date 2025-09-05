import mem
import run
import strings

run(path string) -> error {
	pid := fork()

	if pid == 0 {
		cpath := strings.c(path)
		argv := new(Argv)
		argv.path = cpath.ptr
		err := execve(cpath.ptr, argv, 0)
		mem.free(cpath)
		run.exit(err)
	}

	status := new(int32)
	result := wait4(pid, status, 0, 0)

	if result < 0 {
		return result
	}

	return ([status] >> 8) & 0xFF
}

execve(path *byte, argv *any, envp *any) -> error {
	return syscall(_execve, path, argv, envp)
}

wait4(pid int, status *int32, options int, rusage *any) -> error {
	return syscall(_wait4, pid, status, options, rusage)
}