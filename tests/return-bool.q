main() {
	assert no() == false
	assert yes() == true
	assert notNo() == true
	assert notYes() == false
	assert notNotNo() == false
	assert notNotYes() == true
}

no() -> bool {
	return 1 < 0
}

yes() -> bool {
	return 1 > 0
}

notNo() -> bool {
	return !(1 < 0)
}

notYes() -> bool {
	return !(1 > 0)
}

notNotNo() -> bool {
	return !!(1 < 0)
}

notNotYes() -> bool {
	return !!(1 > 0)
}