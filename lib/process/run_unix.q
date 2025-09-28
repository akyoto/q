import c
import mem
import run

run(path string) -> error {
	pid := fork()

	if pid == 0 {
		cpath := c.string(path)
		argv := new(Argv)
		argv.path = cpath.ptr
		err := execve(cpath.ptr, argv, 0)
		delete(argv)
		mem.free(cpath)
		run.exit(err)
	}

	status := new(int32)
	result := wait4(pid, status, 0, 0)

	if result < 0 {
		delete(status)
		return result
	}

	exitCode := ([status] >> 8) & 0xFF
	delete(status)
	return exitCode
}

execve(path *byte, argv *any, envp *any|nil) -> error {
	return syscall(_execve, path, argv, envp)
}

wait4(pid int, status *int32, options int, rusage *any|nil) -> error {
	return syscall(_wait4, pid, status, options, rusage)
}