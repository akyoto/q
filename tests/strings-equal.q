import strings

main() {
	assert strings.equal("world", "universe") == false
	assert strings.equal("world", "world") == true
	assert strings.equal("世界", "宇宙") == false
	assert strings.equal("世界", "世界") == true
}