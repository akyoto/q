main() {
	sysExit = 20 + 40
	exitCode = sysExit + 15 + sysExit
	syscall(sysExit, exitCode)
}
