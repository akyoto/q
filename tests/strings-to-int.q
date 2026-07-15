import strings

main() {
	_, err0 := strings.toInt("")
	assert err0 != 0

	_, err1 := strings.toInt("abc")
	assert err1 != 0

	n2, err2 := strings.toInt("0")
	assert err2 == 0
	assert n2 == 0

	n3, err3 := strings.toInt("-0")
	assert err3 == 0
	assert n3 == 0

	n4, err4 := strings.toInt("1")
	assert err4 == 0
	assert n4 == 1

	n5, err5 := strings.toInt("-1")
	assert err5 == 0
	assert n5 == -1

	n6, err6 := strings.toInt("65536")
	assert err6 == 0
	assert n6 == 65536

	n7, err7 := strings.toInt("18014398509481984")
	assert err7 == 0
	assert n7 == 18014398509481984
}