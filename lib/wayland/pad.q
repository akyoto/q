pad(length uint) -> uint {
	return (length + 3) & -4
}