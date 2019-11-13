main() {
	sysExit = 60
	exitCode = sysExit + 70
	exitCode = exitCode - 80
	syscall(sysExit, exitCode)
}
